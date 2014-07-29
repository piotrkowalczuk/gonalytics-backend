package controllers

import (
	"net/http"

	"github.com/piotrkowalczuk/gowik-tracker/models"
	"labix.org/v2/mgo/bson"
)

// VisitsGroupedByFirstActionController ...
type VisitsGroupedByFirstActionController struct {
	GeneralController
}

// Get handler
func (vgbfac *VisitsGroupedByFirstActionController) Get() {
	visits := []*models.Visit{}
	timeBucket := vgbfac.Ctx.Input.Param(":timeBucket")

	err := vgbfac.MongoPool.Collection("visit").Find(bson.M{
		"first_action_at_bucket": timeBucket},
	).Select(bson.M{
		"first_action_at": 1,
	}).All(&visits)

	vgbfac.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	vgbfac.ResponseData = models.VisitsGroupedByFirstActionAt(visits)
}
