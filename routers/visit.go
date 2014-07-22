package routers

import (
	"github.com/astaxie/beego"
	"github.com/piotrkowalczuk/gowik-tracker/controllers"
)

func init() {
	beego.Router("/visits", &controllers.VisitsController{})
	beego.Router("/visits/count", &controllers.VisitsCountController{})
}
