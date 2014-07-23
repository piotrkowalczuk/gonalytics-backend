package routers

import (
	"github.com/astaxie/beego"
	"github.com/piotrkowalczuk/gowik-tracker/controllers"
)

func init() {
	beego.Router("/visits", &controllers.VisitsController{})
	beego.Router("/visits/live", &controllers.VisitsLiveController{})
	beego.Router("/visits/count", &controllers.VisitsCountController{})

	beego.Router("/visits/actions", &controllers.VisitsActionsController{})
	beego.Router("/visits/actions/count", &controllers.VisistsActionsCountController{})

	beego.Router("/visits/countries/count", &controllers.VisitsCountriesCountController{})
}
