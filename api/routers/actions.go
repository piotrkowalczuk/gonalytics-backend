package routers

import (
	"github.com/astaxie/beego"
	CtrlV1 "github.com/piotrkowalczuk/gonalytics-backend/api/controllers/v1"
)

func init() {
	beego.Router("/actions", &CtrlV1.ActionsController{})
	beego.Router("/actions/count", &CtrlV1.ActionsCountController{})
}
