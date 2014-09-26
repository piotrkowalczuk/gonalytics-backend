package v1

import "net/http"

// VisitsDistributionController ...
type VisitsDistributionController struct {
	GeneralController
}

// Get ...
func (vdc *VisitsDistributionController) Get() {
	var err error
	var distribution interface{}
	dateTimeRange := vdc.GetString("dateTimeRange")
	by := vdc.GetString("by")

	switch by {
	default:
		distribution, err = vdc.RepositoryManager.Visit.DistributionByTime(dateTimeRange)
	case "country":
		distribution, err = vdc.RepositoryManager.Visit.DistributionByCountry(dateTimeRange)
	}

	vdc.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	vdc.ResponseData = distribution
}
