package service

import (
	"github.com/astaxie/beego/logs"
)

// Instance of Beego logger
var Logger *logs.BeeLogger

// InitLogger ...
func InitLogger() {
	Logger = logs.NewLogger(10000)
	Logger.SetLogger("console", "")
}
