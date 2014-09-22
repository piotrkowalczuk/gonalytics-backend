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
	visits := models.Visits{}
	dateTimeRange := ac.GetString("dateTimeRange")

	query := bson.M{}

	if dateTimeRange != "" {
		query["first_action_at_bucket"] = dateTimeRange
	}

	err := ac.RepositoryManager.Visit.
		Find(query).
		Select(ac.GetQuerySelect()).
		Skip(ac.GetQuerySkip()).
		Limit(ac.GetQueryLimit()).
		All(&visits)

	ac.AbortIf(err, "Unexpected error", http.StatusInternalServerError)
	ac.ResponseData = &visits
}
