package repositories

import "time"

const (
	// SiteMonthCountryActionsCounterColumnFamily ...
	SiteMonthCountryActionsCounterColumnFamily = "site_month_country_actions_counter"
)

// SiteMonthCountryActionsCounterRepository ...
type SiteMonthCountryActionsCounterRepository struct {
	Repository
}

// Increment ...
func (smcacr *SiteMonthCountryActionsCounterRepository) Increment(siteID int64, countryName string, countryCode string, countryID uint, date time.Time) error {

	cql := `
    UPDATE ` + SiteMonthCountryActionsCounterColumnFamily + `
    SET nb_of_actions = nb_of_actions + 1
    WHERE site_id = ?
    AND location_country_name = ?
    AND location_country_code = ?
    AND location_country_id = ?
    AND made_at_year = ?
    AND made_at_month = ?
    `
	return smcacr.Repository.
		Cassandra.
		Query(cql, siteID, countryName, countryCode, countryID, date.Year(), date.Month()).
		Exec()
}
