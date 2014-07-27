package routers

import (
	"github.com/astaxie/beego"
	"github.com/piotrkowalczuk/gowik-tracker/controllers"
)

func init() {
	beego.Router("/visits", &controllers.VisitsController{})
	beego.Router("/visits/live", &controllers.VisitsLiveController{})
	beego.Router("/visits/count", &controllers.VisitsCountController{})
	beego.Router("/visits/group/first-action/:timeBucket", &controllers.VisitsGroupedByFirstActionController{})
	beego.Router("/visits/group/country-code/:timeBucket", &controllers.VisitsGroupedByCountryCodeController{})
	beego.Router("/visits/average-duration", &controllers.VisitsAverageTimeController{})

	beego.Router("/visits/actions", &controllers.VisitsActionsController{})
	beego.Router("/visits/actions/count", &controllers.VisitsActionsCountController{})

	beego.Router("/visits/countries/count", &controllers.VisitsCountriesCountController{})
}
