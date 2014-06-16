package routers

import (
	"github.com/astaxie/beego"
	"github.com/piotrkowalczuk/gonalytics/controllers"
)

func init() {
	beego.Router("/track", &controllers.TrackController{})
}
