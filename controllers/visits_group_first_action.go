package controllers

import (
	"labix.org/v2/mgo/bson"
	"github.com/piotrkowalczuk/gowik-tracker/models"
	"net/http"
)

// VisitsNumberCountController ...
type VisitsGroupedByFirstActionController struct {
	BaseController
}

// Get handler
func (this *VisitsGroupedByFirstActionController) Get() {
	visits := []*models.Visit{}
	timeBucket := this.Ctx.Input.Param(":timeBucket")

	err := this.MongoPool.Collection("visit").Find(bson.M{
		"first_action_at_bucket": timeBucket},
	).Select(bson.M{
		"first_action_at": 1,
	}).All(&visits)

	this.abortIf(err, http.StatusInternalServerError)
	this.Data["json"] = models.VisitsGroupedByFirstActionAt(visits)
	this.ServeJson()
}
