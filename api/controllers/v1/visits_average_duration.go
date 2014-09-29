package v1

import (
	"net/http"

	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
	"labix.org/v2/mgo/bson"
)

// VisitsAverageTimeController ...
type VisitsAverageTimeController struct {
	GeneralController
}

// Get ...
func (vatc *VisitsAverageTimeController) Get() {
	var visits models.Visits
	dateTimeRange := vatc.GetString("dateTimeRange")
	query := bson.M{}

	if dateTimeRange != "" {
		query["first_action_at_bucket"] = dateTimeRange
	}

	err := vatc.MongoDB.C("visit").Find(query).Select(bson.M{
		"first_action_at": 1,
		"last_action_at":  1,
	}).All(&visits)

	vatc.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	vatc.ResponseData = visits.VisitsAverageDuration()
}