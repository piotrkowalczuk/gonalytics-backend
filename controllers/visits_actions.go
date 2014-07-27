package controllers

import (
	"github.com/piotrkowalczuk/gowik-tracker/models"
	"labix.org/v2/mgo/bson"
	"net/http"
)

type VisitsActionsController struct {
	BaseController
}

func (vac *VisitsActionsController) Get() {
	actions := []*models.Action{}
	err := vac.MongoPool.Collection("visits").Find(bson.M{}).All(&actions)

	vac.abortIf(err, http.StatusInternalServerError)
	vac.Data["json"] = &actions
	vac.ServeJson()
}
