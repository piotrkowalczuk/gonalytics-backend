package services

import (
	"github.com/Shopify/sarama"

	"github.com/piotrkowalczuk/gonalytics-backend/lib"
)

// KafkaPublisher ...
var KafkaPublisher *lib.KafkaPublisher

// KafkaClient ...
var KafkaClient *sarama.Client

// InitKafkaClient ...
func InitKafkaClient(name string, config lib.KafkaConfig) {
	var err error
	KafkaClient, err = sarama.NewClient(
		name,
		[]string{config.ConnectionString},
		sarama.NewClientConfig(),
	)

	if err != nil {
		Logger.Error("Connection to Kafka failed.")
		panic(err)
	}

	Logger.Info("Connection to Kafka established sucessfully.")
}

// InitKafkaPublisher ...
func InitKafkaPublisher(config lib.KafkaConfig) {
	KafkaPublisher = &lib.KafkaPublisher{
		Client: KafkaClient,
		Config: config,
	}

	Logger.Info("Kafka publisher initializated successfully.")
}
