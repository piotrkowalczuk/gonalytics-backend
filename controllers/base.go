package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/piotrkowalczuk/gowik-tracker/services"
	"strconv"
)

type (
	BaseController struct {
		beego.Controller
		MongoPool services.Pool
		log       *logs.BeeLogger
	}
)

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

func (bc *BaseController) abortIf(err error, statusCode int) {
	if err != nil {
		bc.log.Error(err.Error())
		bc.Abort(strconv.FormatInt(int64(statusCode), 10))
	}
}
