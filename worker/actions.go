package worker

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/Sirupsen/logrus"
	"github.com/gocql/gocql"
	geoip2 "github.com/oschwald/geoip2-golang"
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
)

type ActionsWorker struct {
	Cassandra         *gocql.Session
	Config            *lib.ActionsWorkerConfig
	KafkaClient       *sarama.Client
	GeoIP             *geoip2.Reader
	Logger            *logrus.Logger
	RepositoryManager lib.RepositoryManager
}

func (aw *ActionsWorker) Start() {
	consumer, err := sarama.NewConsumer(
		aw.KafkaClient,
		"actions",
		0,
		"actions-worker",
		sarama.NewConsumerConfig(),
	)

	if err != nil {
		aw.Logger.Error("Kafka consumer initialization failed.")
		panic(err)
	}

	defer consumer.Close()
	aw.Logger.Info("Actions worker initialized successfully.")

	for {
		select {
		case event := <-consumer.Events():
			if event.Err != nil {
				panic(event.Err)
			}

			trackRequest := &models.TrackRequest{}
			err := json.Unmarshal(event.Value, trackRequest)

			if err != nil {
				aw.Logger.Error(err)
				return
			}

			err = aw.saveAction(trackRequest)
			if err != nil {
				aw.Logger.Error(err)
				return
			}

			aw.Logger.WithFields(logrus.Fields{
				"url":     trackRequest.PageURL,
				"visitId": trackRequest.VisitID,
			}).Debug("Action has been properly processed.")
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
