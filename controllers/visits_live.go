package controllers

import (
	"net/http"
	"time"

	"github.com/piotrkowalczuk/gowik-tracker/models"
	"labix.org/v2/mgo/bson"
)

// VisitsLiveController ...
type VisitsLiveController struct {
	GeneralController
}

// Get ...
func (vlc *VisitsLiveController) Get() {
	limit, err := vlc.GetInt("limit")
	vlc.AbortIf(err, "Missing limit parameter.", http.StatusBadRequest)

	visits := []*models.Visit{}
	err = vlc.MongoPool.Collection("visit").Find(bson.M{
		"last_action_at": bson.M{"$gt": time.Now().Add(-models.MinVisitDuration)},
	}).Sort("-last_action_at").Limit(int(limit)).All(&visits)

	vlc.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	vlc.ResponseData = &visits
}
