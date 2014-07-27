package controllers

import (
	"labix.org/v2/mgo/bson"
	"github.com/piotrkowalczuk/gowik-tracker/models"
	"net/http"
)

// VisitsNumberCountController ...
type VisitsGroupedByCountryCodeController struct {
	BaseController
}

// Get handler
func (this *VisitsGroupedByCountryCodeController) Get() {
	visits := []*models.Visit{}
	timeBucket := this.Ctx.Input.Param(":timeBucket")

	err := this.MongoPool.Collection("visit").Find(bson.M{
		"first_action_at_bucket": timeBucket},
	).Select(bson.M{
		"location": 1,
		"first_action_at": 1,
	}).All(&visits)

	this.abortIf(err, http.StatusInternalServerError)
	this.Data["json"] = models.VisitsGroupedLocationCountryCode(visits)
	this.ServeJson()
}
