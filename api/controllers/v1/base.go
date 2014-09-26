package v1

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/services"
	"labix.org/v2/mgo"
)

// BaseController contains common properties accross multiple controllerss
type BaseController struct {
	beego.Controller
	RepositoryManager lib.RepositoryManager
	MongoDB           *mgo.Database
	Log               *logs.BeeLogger
	Response          interface{}
}

// Prepare is called prior to the basecontrollers method
func (bc *BaseController) Prepare() {
	bc.Log = logs.NewLogger(10000)
	bc.Log.SetLogger("console", "")
	bc.Log.Trace("Http request")
	bc.RepositoryManager = services.RepositoryManager
	bc.MongoDB = services.MongoDB
}

// Finish is called once the basecontrollers method completes
func (bc *BaseController) Finish() {
}

// AbortIf return response if only err is not nil.
func (bc *BaseController) AbortIf(err error, message string, statusCode int) {
	if err != nil {
		bc.Log.Error(err.Error())
		bc.Ctx.Abort(statusCode, message)
	}
}
