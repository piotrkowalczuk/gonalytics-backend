package v1

import (
	"net/http"

	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
)

// VisitsLiveController ...
type VisitsController struct {
	BaseListController
}

// Get ...
func (vlc *VisitsController) Get() {
	visits := models.Visits{}
	actions := models.Actions{}

	visits, err := vlc.RepositoryManager.Visit.Find()

	vlc.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	vlc.ResponseData = &visits
}
