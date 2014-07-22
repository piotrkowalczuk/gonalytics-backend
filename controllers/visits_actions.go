package controllers

import (
	"github.com/piotrkowalczuk/gowik-tracker/models"
	"labix.org/v2/mgo/bson"
)

type VisitsActionsController struct {
	BaseController
}

func (vac *VisitsActionsController) Get() {
	actions := []*models.Action{}
	err := vac.MongoPool.Collection("visits").Find(bson.M{}).All(&actions)

	if err != nil {
		vac.Abort("500")
	}

	vac.Data["json"] = &actions
	vac.ServeJson()
}
