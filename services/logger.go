package services

import (
	"github.com/astaxie/beego/logs"
)

// Singleton instance of logger
var Logger *logs.BeeLogger

// InitLogger ...
func InitLogger() {
	Logger = logs.NewLogger(10000)
	Logger.SetLogger("console", "")
}
