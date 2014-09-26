package v1

import (
	"net/http"
	"time"

	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
	"labix.org/v2/mgo/bson"
)

// VisitsLiveController ...
type VisitsLiveController struct {
	GeneralController
}

// Get ...
func (vlc *VisitsLiveController) Get() {
	visits := models.Visits{}
	actions := models.Actions{}
	limit, err := vlc.GetInt("limit")

	vlc.AbortIf(err, "Missing limit parameter.", http.StatusBadRequest)

	err = vlc.RepositoryManager.Visit.Find(bson.M{
		"last_action_at": bson.M{"$gt": time.Now().Add(-models.MinVisitDuration)},
	}).Sort("-last_action_at").Limit(int(limit)).All(&visits)

	vlc.RepositoryManager.Action.Find(bson.M{
		"_visitId": bson.M{"$in": visits.GetIDs()},
	}).All(&actions)

	for _, action := range actions {
		visit, err := visits.GetByID(action.VisitID)
		if err == nil {
			visit.Actions.Append(action)
		}
	}

	vlc.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	vlc.ResponseData = &visits
}
