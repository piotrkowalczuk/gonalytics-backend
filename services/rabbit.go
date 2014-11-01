package services

import (
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
	"github.com/streadway/amqp"
)

// RabbitMQ ...
var RabbitMQ *amqp.Connection

// InitRabbitMQ ...
func InitRabbitMQ(config lib.RabbitMQConfig) {
	connection, err := amqp.Dial(config.ConnectionString)

	if err != nil {
		Logger.Error("Connection to RabbitMQ failed.")
		panic(err)
	}

	Logger.Info("Connection do RabbitMQ established successfully.")

	RabbitMQ = connection
}
