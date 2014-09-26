package routers

import (
	"github.com/astaxie/beego"
	CtrlV1 "github.com/piotrkowalczuk/gonalytics-backend/tracker/controllers/v1"
)

func init() {
	beego.Router("/track", &CtrlV1.TrackController{})
}
