package repositories

import "time"

const (
	// SiteYearBrowserActionsCounterColumnFamily ...
	SiteYearBrowserActionsCounterColumnFamily = "site_year_browser_actions_counter"
)

// SiteYearBrowserActionsCounterRepository ...
type SiteYearBrowserActionsCounterRepository struct {
	Repository
}

// Increment ...
func (sybacr *SiteYearBrowserActionsCounterRepository) Increment(
	siteID int64,
	browserName string,
	browserVersion string,
	date time.Time,
) error {
	cql := `
    UPDATE ` + SiteYearBrowserActionsCounterColumnFamily + `
    SET nb_of_actions = nb_of_actions + 1
    WHERE site_id = ?
    AND browser_name = ?
    AND browser_version = ?
    AND made_at_year = ?
    `
	return sybacr.Repository.
		Cassandra.
		Query(cql, siteID, browserName, browserVersion, date.Year()).
		Exec()
}
