package controllers

import (
	"github.com/piotrkowalczuk/gowik-tracker/models"
	"labix.org/v2/mgo/bson"
)

type ActionsController struct {
	BaseController
}

func (ac *ActionsController) Get() {
	actions := []*models.Action{}
	err := ac.MongoPool.Collection("visits").Find(bson.M{}).All(&actions)

	if err != nil {
		ac.Abort("500")
	}

	ac.Data["json"] = &actions
	ac.ServeJson()
}
