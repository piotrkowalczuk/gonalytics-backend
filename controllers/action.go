package controllers

import (
	"github.com/piotrkowalczuk/gowik-tracker/models"
	"labix.org/v2/mgo/bson"
)

type ActionListController struct {
	BaseController
}

func (alc *ActionListController) Get() {
	//siteId := rlc.Ctx.Input.Param(":siteId")
	actions := []*models.Action{}
	err := alc.MongoPool.Collection("report").Find(bson.M{}).All(&actions)

	if err != nil {
		alc.Abort("500")
	}

	alc.Data["json"] = &actions
	alc.ServeJson()
}