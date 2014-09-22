package routers

import (
	"github.com/astaxie/beego"
	"github.com/piotrkowalczuk/gonalytics-tracker/controllers"
)

func init() {
	beego.Router("/actions", &controllers.ActionsController{})
	beego.Router("/actions/count", &controllers.ActionsCountController{})
}
