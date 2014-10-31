package v1

import (
	"net/http"
	"time"

	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
)

// VisitsLiveController ...
type VisitsLiveController struct {
	BaseListController
}

// Get ...
func (vlc *VisitsLiveController) Get() {
	visits := models.Visits{}
	actions := models.Actions{}
	limit, err := vlc.GetInt("limit")

	vlc.AbortIf(err, "Missing limit parameter.", http.StatusBadRequest)

	err = vlc.RepositoryManager.Visit.Find(cql.M{
		"last_action_at": cql.M{"$gt": time.Now().Add(-models.MinVisitDuration)},
	}).Sort("-last_action_at").Limit(int(limit)).All(&visits)

	vlc.RepositoryManager.Action.Find(cql.M{
		"_visitId": cql.M{"$in": visits.GetIDs()},
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
