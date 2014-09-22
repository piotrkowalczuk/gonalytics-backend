package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/piotrkowalczuk/gonalytics-tracker/lib"
	"github.com/piotrkowalczuk/gonalytics-tracker/services"
	"labix.org/v2/mgo"
)

// BaseController contains common properties accross multiple controllers
type BaseController struct {
	beego.Controller
	RepositoryManager lib.RepositoryManager
	MongoDB           *mgo.Database
	Log               *logs.BeeLogger
	Response          interface{}
}

// Prepare is called prior to the baseController method
func (bc *BaseController) Prepare() {
	bc.Log = logs.NewLogger(10000)
	bc.Log.SetLogger("console", "")
	bc.Log.Trace("Http request")
	bc.RepositoryManager = services.RepositoryManager
	bc.MongoDB = services.MongoDB
}

// Finish is called once the baseController method completes
func (bc *BaseController) Finish() {
}

// AbortIf return response if only err is not nil.
func (bc *BaseController) AbortIf(err error, message string, statusCode int) {
	if err != nil {
		bc.Log.Error(err.Error())
		bc.Ctx.Abort(statusCode, message)
	}
}
