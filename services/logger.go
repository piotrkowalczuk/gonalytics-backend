package services

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
)

// Singleton instance of logger
var Logger *logrus.Logger

// InitLogger ...
func InitLogger(config lib.LoggerConfig) {
	Logger = logrus.New()
	Logger.Level = logrus.DebugLevel

	if config.Adapter == "console" {
		Logger.Formatter = &lib.ConsoleFormatter{}
	} else {
		panic(errors.New("Unknown logger adapter."))
	}
}
