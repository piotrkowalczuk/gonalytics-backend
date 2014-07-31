package controllers

import (
	"net/http"

	"github.com/piotrkowalczuk/gonalytics-tracker/models"
	"labix.org/v2/mgo/bson"
)

// VisitsController ...
type VisitsController struct {
	GeneralController
}

// Get ..
func (ac *VisitsController) Get() {
	visits := []*models.Visit{}
	dateTimeRange := ac.GetString("dateTimeRange")

	query := bson.M{}

	if dateTimeRange != "" {
		query["created_at_bucket"] = dateTimeRange
	}

	err := ac.MongoPool.Collection("visit").
		Find(query).
		Select(ac.GetQuerySelect()).
		Skip(ac.GetQuerySkip()).
		Limit(ac.GetQueryLimit()).
		All(&visits)

	ac.AbortIf(err, "Unexpected error", http.StatusInternalServerError)
	ac.ResponseData = &visits
}