package routers

import (
	"github.com/astaxie/beego"
	"github.com/piotrkowalczuk/gowik-tracker/controllers"
)

func init() {
	beego.Router("/actions", &controllers.ActionListController{})
}
