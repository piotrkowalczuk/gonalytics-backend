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
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
)

const (
	clientName        = "actions-worker-1"
	consumerGroupName = "actions-workers"
	topicName         = "actions"
	brokerHost        = "localhost:9092"
)

// ActionsWorker ...
type ActionsWorker struct {
	Cassandra         *gocql.Session
	Config            *lib.ActionsWorkerConfig
	KafkaClient       *sarama.Client
	GeoIP             *geoip2.Reader
	Logger            *logrus.Logger
	RepositoryManager lib.RepositoryManager

	startingOffset int64
	broker         *sarama.Broker
	consumer       *sarama.Consumer
}

// ConsumedFunc ...
type ConsumedFunc func(trackRequest *models.TrackRequest) error

func (aw *ActionsWorker) panicIf(err error, message string) {
	if err != nil && err != sarama.NoError {
		aw.Logger.Error(message)
		panic(err)
	}
}

func (aw *ActionsWorker) breakIf(err error) {
	if err != nil && err != sarama.NoError {
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
	defer aw.consumer.Close()

	aw.consume(
		aw.saveAction,
	)
}

func (aw *ActionsWorker) initBroker() {
	aw.broker = sarama.NewBroker(brokerHost)

	err := aw.broker.Open(nil)
	aw.panicIf(err, "Kafka broker initialization failed.")
}

func (aw *ActionsWorker) initConsumer() {
	consumerConfig := &sarama.ConsumerConfig{
		DefaultFetchSize: 32768,
		MinFetchSize:     1,
		MaxWaitTime:      250 * time.Millisecond,
		EventBufferSize:  16,
		OffsetMethod:     sarama.OffsetMethodManual,
		OffsetValue:      aw.startingOffset + 1,
	}

	consumer, err := sarama.NewConsumer(
		aw.KafkaClient,
		topicName,
		0,
		consumerGroupName,
		consumerConfig,
	)

	aw.panicIf(err, "Kafka consumer initialization failed.")

	aw.Logger.WithFields(logrus.Fields{
		"offset": aw.startingOffset,
	}).Info("Kafka consumer initialized successfully.")

	aw.consumer = consumer
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
}

func (aw *ActionsWorker) fetchAndSetStartingOffset() {
	offsetFetchRequest := sarama.OffsetFetchRequest{ConsumerGroup: consumerGroupName}
	offsetFetchRequest.AddPartition(topicName, 0)
	offsetFetchResponse, err := aw.broker.FetchOffset(clientName, &offsetFetchRequest)

	aw.panicIf(err, "Kafka broker fetching current offset request failed.")
	aw.breakIf(offsetFetchResponse.Blocks[topicName][0].Err)

	aw.startingOffset = offsetFetchResponse.Blocks[topicName][0].Offset
}

func (aw *ActionsWorker) consume(callbacks ...ConsumedFunc) {
	for {
		select {
		case event := <-aw.consumer.Events():
			if event.Err != nil {
				panic(event.Err)
			}

			trackRequest := &models.TrackRequest{}
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
			} else if response.Errors[topicName][0] != sarama.NoError {
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

func (aw *ActionsWorker) saveAction(trackRequest *models.TrackRequest) error {
	actionCreator := lib.NewActionCreator(aw.GeoIP)
	action, err := actionCreator.Create(trackRequest)
	if err != nil {
		return err
	}

	err = aw.RepositoryManager.Action.Insert(action)
	if err != nil {
		return err
	}

	return nil
}
