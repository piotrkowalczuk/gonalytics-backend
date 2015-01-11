package repositories

import "time"

const (
	// SiteDayBrowserActionsCounterColumnFamily ...
	SiteDayBrowserActionsCounterColumnFamily = "site_day_browser_actions_counter"
)

// SiteDayBrowserActionsCounterRepository ...
type SiteDayBrowserActionsCounterRepository struct {
	Repository
}

// Increment ...
func (sdbacr *SiteDayBrowserActionsCounterRepository) Increment(
	siteID int64,
	browserName string,
	browserVersion string,
	date time.Time,
) error {
	cql := `
    UPDATE ` + SiteDayBrowserActionsCounterColumnFamily + `
    SET nb_of_actions = nb_of_actions + 1
    WHERE site_id = ?
    AND browser_name = ?
    AND browser_version = ?
    AND made_at_year = ?
    AND made_at_month = ?
    AND made_at_day = ?
    `
	return sdbacr.Repository.
		Cassandra.
		Query(cql, siteID, browserName, browserVersion, date.Year(), date.Month(), date.Day()).
		Exec()
}
