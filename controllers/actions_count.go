package controllers

import (
	"labix.org/v2/mgo/bson"
)

type ActionsCountController struct {
	BaseController
}

func (acc *ActionsCountController) Get() {
	dateTimeRange := acc.GetString("dateTimeRange")
	numberOfActions, err := acc.MongoPool.Collection("visit").Find(bson.M{"actions.created_at.bucket": dateTimeRange}).Count()

	if err != nil {
		acc.Abort("500")
	}

	acc.Data["json"] = numberOfActions
	acc.ServeJson()
}
