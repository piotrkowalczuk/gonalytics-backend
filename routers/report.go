package routers

import (
	"github.com/astaxie/beego"
	"github.com/piotrkowalczuk/gonalytics-tracker/controllers"
)

func init() {
	beego.Router("/report/country", &controllers.ReportCountryController{})
}
