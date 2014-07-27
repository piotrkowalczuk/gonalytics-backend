package controllers

import (
	"labix.org/v2/mgo/bson"
	"net/http"
)

type VisitsCountController struct {
	BaseController
}

func (vcc *VisitsCountController) Get() {
	dateTimeRange := vcc.GetString("dateTimeRange")
	numberOfVisits, err := vcc.MongoPool.Collection("visit").Find(
		bson.M{"created_at_bucket": dateTimeRange},
	).Count()

	vcc.abortIf(err, http.StatusInternalServerError)
	vcc.Data["json"] = numberOfVisits
	vcc.ServeJson()
}
