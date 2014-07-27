package controllers

import (
	"github.com/piotrkowalczuk/gowik-tracker/models"
	"labix.org/v2/mgo/bson"
	"net/http"
)

type VisitsAverageTimeController struct {
	BaseController
}

func (vatc *VisitsAverageTimeController) Get() {
	visits := []*models.Visit{}
	dateTimeRange := vatc.GetString("dateTimeRange")
	query := bson.M{}

	if dateTimeRange != "" {
		query["created_at_bucket"] = dateTimeRange
	}

	err := vatc.MongoPool.Collection("visit").Find(query).Select(bson.M{
		"first_action_at": 1,
		"last_action_at":  1,
	}).All(&visits)

	vatc.abortIf(err, http.StatusInternalServerError)
	vatc.Data["json"] = models.VisitsAverageDuration(visits)
	vatc.ServeJson()
}
