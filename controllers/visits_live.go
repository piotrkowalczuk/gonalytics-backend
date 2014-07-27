package controllers

import (
	"github.com/piotrkowalczuk/gowik-tracker/models"
	"labix.org/v2/mgo/bson"
	"net/http"
	"time"
)

type VisitsLiveController struct {
	BaseController
}

func (vlc *VisitsLiveController) Get() {
	limit, err := vlc.GetInt("limit")
	vlc.abortIf(err, http.StatusBadRequest)

	visits := []*models.Visit{}
	err = vlc.MongoPool.Collection("visit").Find(bson.M{
		"last_action_at": bson.M{"$gt": time.Now().Add(-models.MIN_VISIT_DURATION)},
	}).Sort("-last_action_at").Limit(int(limit)).All(&visits)

	vlc.abortIf(err, http.StatusInternalServerError)
	vlc.Data["json"] = &visits
	vlc.ServeJson()
}
