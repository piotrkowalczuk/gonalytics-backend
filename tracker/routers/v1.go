package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/services"
	"github.com/piotrkowalczuk/gonalytics-backend/tracker/controllers/v1"
)

// GetNamespaceV1 ...
func GetNamespaceV1() *beego.Namespace {
	baseController := v1.BaseController{
		Log:               logs.NewLogger(10000),
		RepositoryManager: services.RepositoryManager,
	}

	return beego.NewNamespace("/v1",
		beego.NSRouter("/visit", &v1.VisitController{baseController}),
	)
}
