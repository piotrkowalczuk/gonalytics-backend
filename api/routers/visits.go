package routers

import (
	"github.com/astaxie/beego"
	CtrlV1 "github.com/piotrkowalczuk/gonalytics-backend/api/controllers/v1"
)

func init() {
	beego.Router("/visits", &CtrlV1.VisitsController{})
	beego.Router("/visits/live", &CtrlV1.VisitsLiveController{})
	beego.Router("/visits/count", &CtrlV1.VisitsCountController{})
	beego.Router("/visits/average-duration", &CtrlV1.VisitsAverageTimeController{})
	beego.Router("/visits/distribution", &CtrlV1.VisitsDistributionController{})
}
