package v1

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
	"github.com/piotrkowalczuk/gonalytics-backend/services"
)

// BaseController contains common properties accross multiple controllerss
type BaseController struct {
	beego.Controller
	RepositoryManager lib.RepositoryManager
	Log               *logs.BeeLogger
	ResponseData      interface{}
}

// Prepare is called prior to the basecontrollers method
func (bc *BaseController) Prepare() {
	bc.Log = services.Logger
}

// Finish is called once the basecontrollers method completes
func (bc *BaseController) Finish() {
	bc.Data["json"] = &bc.ResponseData
	bc.ServeJson()
}

// AbortIf return response if only err is not nil.
func (bc *BaseController) AbortIf(err error, message string, statusCode int) {
	if err != nil {
		bc.Log.Error(err.Error())
		bc.Ctx.Abort(statusCode, message)
	}
}
