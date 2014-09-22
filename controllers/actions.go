package controllers

import (
	"net/http"

	"github.com/piotrkowalczuk/gonalytics-tracker/models"
	"labix.org/v2/mgo/bson"
)

// VisitsActionsController ...
type ActionsController struct {
	GeneralController
}

// Get ...
func (ac *ActionsController) Get() {
	actions := []*models.Action{}
	err := ac.RepositoryManager.Action.Find(bson.M{}).All(&actions)

	ac.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	ac.ResponseData = &actions
}
