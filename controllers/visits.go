package controllers

import (
	"github.com/piotrkowalczuk/gowik-tracker/models"
	"labix.org/v2/mgo/bson"
)

type VisitsController struct {
	BaseController
}

func (ac *VisitsController) Get() {
	visits := []*models.Visit{}
	err := ac.MongoPool.Collection("visit").Find(bson.M{}).All(&visits)

	if err != nil {
		ac.Abort("500")
	}

	ac.Data["json"] = &visits
	ac.ServeJson()
}
