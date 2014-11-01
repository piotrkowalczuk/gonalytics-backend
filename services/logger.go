package services

import (
	"github.com/astaxie/beego/logs"
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
)

// Singleton instance of logger
var Logger *logs.BeeLogger

// InitLogger ...
func InitLogger(config lib.LoggerConfig) {
	Logger = logs.NewLogger(config.NbOfChannels)
	Logger.SetLogger(config.Adapter, config.Settings)
}
