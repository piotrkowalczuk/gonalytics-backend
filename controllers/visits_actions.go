package controllers

import (
	"net/http"

	"github.com/piotrkowalczuk/gowik-tracker/models"
	"labix.org/v2/mgo/bson"
)

// VisitsActionsController ...
type VisitsActionsController struct {
	GeneralController
}

// Get ...
func (vac *VisitsActionsController) Get() {
	actions := []*models.Action{}
	err := vac.MongoPool.Collection("visits").Find(bson.M{}).All(&actions)

	vac.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	vac.ResponseData = &actions
}
