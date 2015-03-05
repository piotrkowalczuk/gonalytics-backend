package worker

import (
	"encoding/json"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/Sirupsen/logrus"
	"github.com/gocql/gocql"
	geoip2 "github.com/oschwald/geoip2-golang"
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/repositories"
)

const (
	clientName        = "actions-worker-1"
	consumerGroupName = "actions-workers"
	topicName         = "actions"
)

// ActionsWorker ...
type ActionsWorker struct {
	Cassandra         *gocql.Session
	Config            *lib.ActionsWorkerConfig
	KafkaClient       *sarama.Client
	GeoIP             *geoip2.Reader
	Logger            *logrus.Logger
	RepositoryManager lib.RepositoryManager

	BrokerHost string

	startingOffset    int64
	broker            *sarama.Broker
	consumer          *sarama.Consumer
	partitionConsumer *sarama.PartitionConsumer
}

// ConsumedFunc ...
type ConsumedFunc func(trackRequest *lib.TrackRequest) error

func (aw *ActionsWorker) panicIf(err error, message string) {
	if err != nil && err != sarama.ErrNoError {
		aw.Logger.Error(message)
		panic(err)
	}
}

func (aw *ActionsWorker) breakIf(err error) {
	if err != nil && err != sarama.ErrNoError {
		aw.Logger.Error(err.Error())
		os.Exit(1)
	}
}

// Start ...
func (aw *ActionsWorker) Start() {
	aw.initBroker()
	defer aw.broker.Close()

	aw.logNumberOfTopics()
	aw.checkIfAwailableOffsets()
	aw.fetchAndSetStartingOffset()

	aw.initConsumer()
	defer aw.partitionConsumer.Close()

	aw.consume(
		aw.saveToCassandra,
	)
}

func (aw *ActionsWorker) initBroker() {
	aw.broker = sarama.NewBroker(aw.BrokerHost)

	err := aw.broker.Open(nil)
	aw.panicIf(err, "Kafka broker initialization failed.")
}

func (aw *ActionsWorker) initConsumer() {
	consumerConfig := &sarama.ConsumerConfig{
		MinFetchSize: 1,
		MaxWaitTime:  250 * time.Millisecond,
	}

	partitionConsumerConfig := &sarama.PartitionConsumerConfig{
		DefaultFetchSize: 32768,
		OffsetMethod:     sarama.OffsetMethodManual,
		OffsetValue:      aw.startingOffset + 1,
	}

	consumer, err := sarama.NewConsumer(
		aw.KafkaClient,
		consumerConfig,
	)

	aw.panicIf(err, "Kafka consumer initialization failed.")

	partitionConsumer, err := consumer.ConsumePartition(
		topicName,
		0,
		partitionConsumerConfig,
	)

	aw.panicIf(err, "Kafka partition consumer initialization failed.")

	aw.Logger.WithFields(logrus.Fields{
		"offset": aw.startingOffset,
	}).Info("Kafka consumer initialized successfully.")

	aw.consumer = consumer
	aw.partitionConsumer = partitionConsumer
}

func (aw *ActionsWorker) logNumberOfTopics() {
	request := sarama.MetadataRequest{Topics: []string{}}
	response, err := aw.broker.GetMetadata(clientName, &request)
	aw.breakIf(err)

	aw.Logger.WithFields(logrus.Fields{
		"nbOfTopics": len(response.Topics),
	}).Info("Kafka metadata fetched successfully.")
}

func (aw *ActionsWorker) checkIfAwailableOffsets() {
	offsetsRequest := sarama.OffsetRequest{}
	offsetsRequest.AddBlock(topicName, 0, sarama.EarliestOffset, 255)
	offsetsResponse, err := aw.broker.GetAvailableOffsets(clientName, &offsetsRequest)

	aw.panicIf(err, "Kafka broker available offsets request failed.")
	aw.breakIf(offsetsResponse.Blocks[topicName][0].Err)

	aw.Logger.WithFields(logrus.Fields{
		"offsets": offsetsResponse.Blocks[topicName][0].Offsets,
	}).Info("Kafka available offsets fetched successfully.")
}

func (aw *ActionsWorker) fetchAndSetStartingOffset() {
	offsetFetchRequest := sarama.OffsetFetchRequest{ConsumerGroup: consumerGroupName}
	offsetFetchRequest.AddPartition(topicName, 0)
	offsetFetchResponse, err := aw.broker.FetchOffset(clientName, &offsetFetchRequest)

	aw.panicIf(err, "Kafka broker fetching current offset request failed.")

	err = offsetFetchResponse.Blocks[topicName][0].Err
	if err != sarama.ErrNoError {
		if err == sarama.ErrUnknownTopicOrPartition {
			offsetCommitRequest := &sarama.OffsetCommitRequest{ConsumerGroup: consumerGroupName}
			offsetCommitRequest.AddBlock(topicName, 0, 0, sarama.ReceiveTime, "")

			response, err := aw.broker.CommitOffset(clientName, offsetCommitRequest)

			if err != nil {
				aw.Logger.Error(err)
				return
			} else if response == nil {
				aw.Logger.Error("Brocker returns no response")
				return
			} else if response.Errors[topicName][0] != sarama.ErrNoError {
				aw.Logger.Error(response.Errors[topicName][0])
				return
			}

			aw.startingOffset = 0
		} else {
			aw.breakIf(offsetFetchResponse.Blocks[topicName][0].Err)
		}
	} else {
		aw.startingOffset = offsetFetchResponse.Blocks[topicName][0].Offset
	}

	aw.Logger.WithFields(logrus.Fields{
		"startingOffset": aw.startingOffset,
	}).Info("Kafka current offsets fetched successfully.")
}

func (aw *ActionsWorker) consume(callbacks ...ConsumedFunc) {
	for {
		select {
		case err := <-aw.partitionConsumer.Errors():
			if err != nil {
				panic(err)
			}
		case event := <-aw.partitionConsumer.Messages():
			trackRequest := &lib.TrackRequest{}
			err := json.Unmarshal(event.Value, trackRequest)

			if err != nil {
				aw.Logger.Error(err)
				return
			}

			for _, fn := range callbacks {
				err := fn(trackRequest)
				if err != nil {
					aw.Logger.Error(err)
					return
				}
			}

			aw.Logger.WithFields(logrus.Fields{
				"url":     trackRequest.PageURL,
				"visitId": trackRequest.VisitID,
				"offset":  event.Offset,
			}).Debug("Action has been properly processed.")

			offsetCommitRequest := &sarama.OffsetCommitRequest{ConsumerGroup: consumerGroupName}
			offsetCommitRequest.AddBlock(topicName, 0, event.Offset, sarama.ReceiveTime, "")

			response, err := aw.broker.CommitOffset(clientName, offsetCommitRequest)
			if err != nil {
				aw.Logger.Error(err)
				return
			} else if response == nil {
				aw.Logger.Error("Brocker returns no response")
				return
			} else if response.Errors[topicName][0] != sarama.ErrNoError {
				aw.Logger.Error(response.Errors[topicName][0])
				return
			}

			aw.Logger.WithFields(logrus.Fields{
				"client":        clientName,
				"consumerGroup": consumerGroupName,
				"topic":         topicName,
				"offset":        event.Offset,
			}).Debug("New offset has been successfully submited to brocker.")
		}
	}
}

func (aw *ActionsWorker) saveToCassandra(trackRequest *lib.TrackRequest) error {
	actionCreator := lib.NewActionCreator(aw.GeoIP)
	metricRequestMatcher := lib.NewMetricRequestMatcherFromConfig(aw.Config)

	matchingMetrics := metricRequestMatcher.Matching(trackRequest)

	// Multidimensional metric counters increment
	for _, metric := range matchingMetrics {
		dimensionsNames, dimensionsValues := metric.DimensionsNamesAndDimensionsValues()

		var err error

		now := time.Now().UTC()
		incrementableRepositories := []repositories.MetricIncrementer{
			// Metrics by value
			&aw.RepositoryManager.MetricDayByValueCounter,
			// Metrics by time frame
			&aw.RepositoryManager.MetricDayByMinuteCounter,
			&aw.RepositoryManager.MetricMonthByHourCounter,
			&aw.RepositoryManager.MetricYearByDayCounter,
		}
		incrementor := func(repository repositories.MetricIncrementer) {
			if err != nil {
				return
			}

			err = repository.Increment(dimensionsNames, dimensionsValues, now)
		}

		for _, repository := range incrementableRepositories {
			incrementor(repository)
		}

		if err != nil {
			return err
		}

		aw.Logger.WithFields(logrus.Fields{
			"dimensionsNames":  dimensionsNames,
			"dimensionsValues": dimensionsValues,
		}).Debug("Metric counters has been successfully incremented.")
	}

	action, err := actionCreator.Create(trackRequest)
	if err != nil {
		return err
	}

	return aw.RepositoryManager.VisitActions.Insert(action)
}
