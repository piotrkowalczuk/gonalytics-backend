package lib

import (
	"github.com/Shopify/sarama"
)

// KafkaPublisher ...
type KafkaPublisher struct {
	Client *sarama.Client
	Config KafkaConfig
}

// PublishAction ...
func (kp *KafkaPublisher) PublishAction(message string) error {
	producer, err := sarama.NewProducer(kp.Client, nil)
	if err != nil {
		return err
	}
	defer producer.Close()

	select {
	case producer.Input() <- &sarama.ProducerMessage{
		Topic: "actions",
		Key:   nil,
		Value: sarama.StringEncoder(message),
	}:
		return nil
	case err := <-producer.Errors():
		return err.Err
	}
}
