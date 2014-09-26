package v1

import (
	"net/http"

	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
	"labix.org/v2/mgo/bson"
)

// ActionsController ...
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
