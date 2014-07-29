package controllers

import (
	"net/http"

	"labix.org/v2/mgo/bson"
)

// VisitsCountController ...
type VisitsCountController struct {
	GeneralController
}

// Get ...
func (vcc *VisitsCountController) Get() {
	dateTimeRange := vcc.GetString("dateTimeRange")
	numberOfVisits, err := vcc.MongoPool.Collection("visit").Find(
		bson.M{"first_action_at_bucket": dateTimeRange},
	).Count()

	vcc.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	vcc.ResponseData = numberOfVisits
}
