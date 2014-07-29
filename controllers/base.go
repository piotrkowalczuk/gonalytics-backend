package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/piotrkowalczuk/gowik-tracker/services"
)

// BaseController contains common properties accross multiple controllers
type BaseController struct {
	beego.Controller
	MongoPool services.Pool
	log       *logs.BeeLogger
	response  interface{}
}

// Prepare is called prior to the baseController method
func (bc *BaseController) Prepare() {
	bc.log = logs.NewLogger(10000)
	bc.log.SetLogger("console", "")
	bc.log.Trace("Http request")
	bc.MongoPool = services.MongoPool
}

// Finish is called once the baseController method completes
func (bc *BaseController) Finish() {
}

// AbortIf return response if only err is not nil.
func (bc *BaseController) AbortIf(err error, message string, statusCode int) {
	if err != nil {
		bc.log.Error(err.Error())
		bc.Ctx.Abort(statusCode, message)
	}
}
