package controllers

import (
	"labix.org/v2/mgo/bson"
)

type ActionCountController struct {
	BaseController
}

func (acc *ActionCountController) Get() {
	dateTimeRange := acc.GetString("dateTimeRange")
	numberOfActions, err := acc.MongoPool.Collection("action").Find(bson.M{"created_at.bucket": dateTimeRange}).Count()

	if err != nil {
		acc.Abort("500")
	}

	acc.Data["json"] = numberOfActions
	acc.ServeJson()
}
