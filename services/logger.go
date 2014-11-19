package services

import (
	"github.com/Sirupsen/logrus"
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
	// "os"
)

// Singleton instance of logger
var Logger *logrus.Logger

// InitLogger ...
func InitLogger(config lib.LoggerConfig) {
	Logger = logrus.New()
	Logger.Formatter = &lib.ConsoleFormatter{}
}
