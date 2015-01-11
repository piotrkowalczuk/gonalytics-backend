package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gocraft/web"
)

// NbOfVisitsByCountryGETHandler ...
func (bc *BaseContext) NbOfVisitsByCountryGETHandler(w web.ResponseWriter, r *web.Request) {
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
		result, err = bc.RepositoryManager.SiteDayCountryVisitsCounter.Find(siteID, date)
	case AggregateByMonth:
		result, err = bc.RepositoryManager.SiteMonthCountryVisitsCounter.Find(siteID, date)
	case AggregateByYear:
		result, err = bc.RepositoryManager.SiteYearCountryVisitsCounter.Find(siteID, date)
	default:
		result, err = bc.RepositoryManager.SiteDayCountryVisitsCounter.Find(siteID, date)
	}

	if err != nil {
		bc.HTTPError(w, err, "Unexpected error.", http.StatusInternalServerError)
		return
	}

	bc.ServeJSON(w, result)
}
