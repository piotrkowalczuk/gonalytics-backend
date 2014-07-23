package controllers

import (
	"labix.org/v2/mgo/bson"
)

type VisitsCountController struct {
	BaseController
}

func (vcc *VisitsCountController) Get() {
	dateTimeRange := vcc.GetString("dateTimeRange")
	numberOfVisits, err := vcc.MongoPool.Collection("visit").Find(
		bson.M{"created_at_bucket": dateTimeRange},
	).Count()

	if err != nil {
		vcc.Abort("500")
	}

	vcc.Data["json"] = numberOfVisits
	vcc.ServeJson()
}
