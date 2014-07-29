package controllers

import (
	"net/http"

	"github.com/piotrkowalczuk/gowik-tracker/models"
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
	err := vcc.MongoPool.Collection("visit").Find(
		bson.M{"created_at_bucket": dateTimeRange},
	).All(&visits)

	vcc.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	vcc.ResponseData = models.NewCountryReportFromVisits(visits)
}
