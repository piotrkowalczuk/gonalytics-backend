package routers

import (
	"github.com/astaxie/beego"
	CtrlV1 "github.com/piotrkowalczuk/gonalytics-backend/api/controllers/v1"
)

func init() {
	beego.Router("/report/country", &CtrlV1.CountryReportController{})
}
