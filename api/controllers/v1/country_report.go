package v1

import (
	"net/http"

	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/reports"
	"labix.org/v2/mgo/bson"
)

// CountryReportController ...
type CountryReportController struct {
	GeneralController
}

// Get ...
func (crc *CountryReportController) Get() {
	dateTimeRange := crc.GetString("dateTimeRange")
	visits := []*models.Visit{}
	err := crc.RepositoryManager.Visit.Find(
		bson.M{"first_action_at_bucket": dateTimeRange},
	).All(&visits)

	crc.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	crc.ResponseData = reports.NewCountryReportFromVisits(visits)
}
