package controllers

import (
	"github.com/piotrkowalczuk/gowik-tracker/models"
	"labix.org/v2/mgo/bson"
	"net/http"
)

type ReportCountryController struct {
	BaseController
}

func (vcc *ReportCountryController) Get() {
	dateTimeRange := vcc.GetString("dateTimeRange")
	visits := []*models.Visit{}
	err := vcc.MongoPool.Collection("visit").Find(
		bson.M{"created_at_bucket": dateTimeRange},
	).All(&visits)

	vcc.abortIf(err, http.StatusInternalServerError)
	vcc.Data["json"] = models.NewCountryReportFromVisits(visits)
	vcc.ServeJson()
}
