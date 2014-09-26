package v1

import (
	"net/http"

	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
	"labix.org/v2/mgo/bson"
)

// VisitsController ...
type VisitsController struct {
	GeneralController
}

// Get ..
func (vc *VisitsController) Get() {
	visits := models.Visits{}
	dateTimeRange := vc.GetString("dateTimeRange")

	query := bson.M{}

	if dateTimeRange != "" {
		query["first_action_at_bucket"] = dateTimeRange
	}

	err := vc.RepositoryManager.Visit.
		Find(query).
		Select(vc.GetQuerySelect()).
		Skip(vc.GetQuerySkip()).
		Limit(vc.GetQueryLimit()).
		All(&visits)

	vc.AbortIf(err, "Unexpected error", http.StatusInternalServerError)
	vc.ResponseData = &visits
}
