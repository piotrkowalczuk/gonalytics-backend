package controllers

import (
	"net/http"

	"github.com/piotrkowalczuk/gonalytics-tracker/models"
	"labix.org/v2/mgo/bson"
)

// ReportCountryController ...
type ReportCountryController struct {
	GeneralController
}

// Get ...
func (vcc *ReportCountryController) Get() {
	dateTimeRange := vcc.GetString("dateTimeRange")
	visits := []*models.Visit{}
	err := vcc.RepositoryManager.Visit.Find(
		bson.M{"first_action_at_bucket": dateTimeRange},
	).All(&visits)

	vcc.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	vcc.ResponseData = models.NewCountryReportFromVisits(visits)
}
