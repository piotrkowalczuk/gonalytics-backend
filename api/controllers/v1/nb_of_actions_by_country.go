package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gocraft/web"
)

// NbOfActionsByCountryGETHandler ...
func (bc *BaseContext) NbOfActionsByCountryGETHandler(w web.ResponseWriter, r *web.Request) {
	r.ParseForm()

	siteID, err := strconv.ParseInt(r.PathParams["siteId"], 10, 64)
	if err != nil {
		bc.HTTPError(w, err, "Unexpected error.", http.StatusBadRequest)
		return
	}

	date := time.Now()
	if timestampString := r.FormValue("timestamp"); timestampString != "" {
		timestamp, err := strconv.ParseInt(timestampString, 10, 64)
		if err != nil {
			bc.HTTPError(w, err, "Unexpected error.", http.StatusBadRequest)
			return
		}

		date = time.Unix(timestamp, 0)
	}

	var result interface{}
	switch r.FormValue("aggregationType") {
	case AggregateByDay:
		result, err = bc.RepositoryManager.SiteDayCountryActionsCounter.Find(siteID, date)
	case AggregateByMonth:
		result, err = bc.RepositoryManager.SiteMonthCountryActionsCounter.Find(siteID, date)
	case AggregateByYear:
		result, err = bc.RepositoryManager.SiteYearCountryActionsCounter.Find(siteID, date)
	default:
		result, err = bc.RepositoryManager.SiteDayCountryActionsCounter.Find(siteID, date)
	}

	if err != nil {
		bc.HTTPError(w, err, "Unexpected error.", http.StatusInternalServerError)
		return
	}

	bc.ServeJSON(w, result)
}
