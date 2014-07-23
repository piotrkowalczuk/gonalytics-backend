package controllers

import (
	"github.com/piotrkowalczuk/gowik-tracker/models"
	"labix.org/v2/mgo/bson"
	"time"
)

type VisitsLiveController struct {
	BaseController
}

func (vlc *VisitsLiveController) Get() {
	visits := []*models.Visit{}
	err := vlc.MongoPool.Collection("visit").Find(bson.M{
		"last_action_at": bson.M{"$gt": time.Now().Add(-30 * time.Minute)},
	}).Sort("-last_action_at").All(&visits)

	if err != nil {
		vlc.Abort("500")
	}

	vlc.Data["json"] = &visits
	vlc.ServeJson()
}
