package routers

import (
	"github.com/astaxie/beego"
	"github.com/piotrkowalczuk/gonalytics-backend/tracker/controllers/v1"
)

// GetNamespaceV1 ...
func GetNamespaceV1() *beego.Namespace {
	return beego.NewNamespace("/v1",
		beego.NSRouter("/visit", &v1.VisitController{}),
	)
}
