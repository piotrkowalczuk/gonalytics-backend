package controllers

import (
	"net/http"

	"labix.org/v2/mgo/bson"
)

// VisitsCountController ...
type VisitsCountController struct {
	GeneralController
	ResponseData struct {
		NumberOfVisits int `json:"nbOfVisits"`
	}
}

// Get ...
func (vcc *VisitsCountController) Get() {
	var err error

	dateTimeRange := vcc.GetString("dateTimeRange")
	vcc.ResponseData.NumberOfVisits, err = vcc.MongoPool.Collection("visit").Find(
		bson.M{"first_action_at_bucket": dateTimeRange},
	).Count()

	vcc.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
}
