package controllers

import (
	"github.com/piotrkowalczuk/gowik-tracker/models"
	"labix.org/v2/mgo/bson"
)

type ReportCountryController struct {
	BaseController
}

func (vcc *ReportCountryController) Get() {
	dateTimeRange := vcc.GetString("dateTimeRange")
	visits := []*models.Visit{}
	err := vcc.MongoPool.Collection("visit").Find(bson.M{"created_at.bucket": dateTimeRange}).All(&visits)

	if err != nil {
		vcc.Abort("500")

	}

	vcc.Data["json"] = models.NewCountryReportFromVisits(visits)
	vcc.ServeJson()
}
