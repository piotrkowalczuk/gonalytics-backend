package controllers

import "net/http"

// VisitsCountController ...
type VisitsCountController struct {
	GeneralController
}

// VisitsCountResponse ...
type VisitsCountResponse struct {
	NumberOfVisits int64 `json:"nbOfVisits"`
}

// Get ...
func (vcc *VisitsCountController) Get() {
	var err error
	var queryFunc func(dateTimeRange string) (int64, error)

	response := VisitsCountResponse{
		NumberOfVisits: 0,
	}

	dateTimeRange := vcc.GetString("dateTimeRange")
	groupBy := vcc.GetString("groupBy")

	switch groupBy {
	default:
		queryFunc = vcc.RepositoryManager.Visit.Count
	case "country":
		queryFunc = vcc.RepositoryManager.Visit.CountByCountryID
	}

	response.NumberOfVisits, err = queryFunc(dateTimeRange)

	vcc.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	vcc.ResponseData = response
}
