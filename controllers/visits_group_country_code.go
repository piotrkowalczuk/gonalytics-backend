package controllers

import (
	"net/http"

	"github.com/piotrkowalczuk/gowik-tracker/models"
	"labix.org/v2/mgo/bson"
)

// VisitsGroupedByCountryCodeController ...
type VisitsGroupedByCountryCodeController struct {
	GeneralController
}

// Get handler
func (vgbccc *VisitsGroupedByCountryCodeController) Get() {
	visits := []*models.Visit{}
	timeBucket := vgbccc.Ctx.Input.Param(":timeBucket")

	err := vgbccc.MongoPool.Collection("visit").Find(bson.M{
		"first_action_at_bucket": timeBucket},
	).Select(bson.M{
		"location":        1,
		"first_action_at": 1,
	}).All(&visits)

	vgbccc.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	vgbccc.ResponseData = models.VisitsGroupedLocationCountryCode(visits)
}
