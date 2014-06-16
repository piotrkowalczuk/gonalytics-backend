package controllers

import (
	"github.com/piotrkowalczuk/gonalytics/models"
	// "labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type ReportListController struct {
	BaseController
}

func (rlc *ReportListController) Get() {
	//siteId := rlc.Ctx.Input.Param(":siteId")
	reports := []*models.Report{}
	err := rlc.MongoPool.Collection("report").Find(bson.M{}).All(&reports)

	if err != nil {
		rlc.Abort("500")
	}

	rlc.Data["json"] = &reports
	rlc.ServeJson()
}
