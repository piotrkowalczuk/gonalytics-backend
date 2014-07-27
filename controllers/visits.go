package controllers

import (
	"github.com/piotrkowalczuk/gowik-tracker/models"
	"labix.org/v2/mgo/bson"
	"net/http"
)

type VisitsController struct {
	BaseController
}

func (ac *VisitsController) Get() {
	visits := []*models.Visit{}
	dateTimeRange := ac.GetString("dateTimeRange")

	query := bson.M{}

	if dateTimeRange != "" {
		query["created_at_bucket"] = dateTimeRange
	}

	err := ac.MongoPool.Collection("visit").Find(query).All(&visits)

	ac.abortIf(err, http.StatusInternalServerError)
	ac.Data["json"] = &visits
	ac.ServeJson()
}
