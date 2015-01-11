package repositories

import "time"

const (
	// SiteMonthBrowserActionsCounterColumnFamily ...
	SiteMonthBrowserActionsCounterColumnFamily = "site_month_browser_actions_counter"
)

// SiteMonthBrowserActionsCounterRepository ...
type SiteMonthBrowserActionsCounterRepository struct {
	Repository
}

// Increment ...
func (smbacr *SiteMonthBrowserActionsCounterRepository) Increment(
	siteID int64,
	browserName string,
	browserVersion string,
	date time.Time,
) error {
	cql := `
    UPDATE ` + SiteMonthBrowserActionsCounterColumnFamily + `
    SET nb_of_actions = nb_of_actions + 1
    WHERE site_id = ?
    AND browser_name = ?
    AND browser_version = ?
    AND made_at_year = ?
    AND made_at_month = ?
    `
	return smbacr.Repository.
		Cassandra.
		Query(cql, siteID, browserName, browserVersion, date.Year(), date.Month()).
		Exec()
}
