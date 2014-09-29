package routers

import (
	"github.com/astaxie/beego"
	"github.com/piotrkowalczuk/gonalytics-backend/api/controllers/v1"
)

// InitNamespaceV1 ...
func InitNamespaceV1() *beego.Namespace {
	return beego.NewNamespace("/v1",
		beego.NSRouter("/visits", &v1.VisitsController{}),
		beego.NSRouter("/visits/live", &v1.VisitsLiveController{}),
		beego.NSRouter("/visits/count", &v1.VisitsCountController{}),
		beego.NSRouter("/visits/average-duration", &v1.VisitsAverageTimeController{}),
		beego.NSRouter("/visits/distribution", &v1.VisitsDistributionController{}),
		beego.NSRouter("/actions", &v1.ActionsController{}),
		beego.NSRouter("/actions/count", &v1.ActionsCountController{}),
		beego.NSRouter("/report/country", &v1.CountryReportController{}),
	)
}
