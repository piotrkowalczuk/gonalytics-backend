package repositories

import "time"

const (
	// SiteDayCountryActionsCounterColumnFamily ...
	SiteDayCountryActionsCounterColumnFamily = "site_day_country_actions_counter"
)

// SiteDayCountryActionsCounterRepository ...
type SiteDayCountryActionsCounterRepository struct {
	Repository
}

// Increment ...
func (sdcacr *SiteDayCountryActionsCounterRepository) Increment(siteID int64, countryName string, countryCode string, countryID uint, date time.Time) error {

	cql := `
    UPDATE ` + SiteDayCountryActionsCounterColumnFamily + `
    SET nb_of_actions = nb_of_actions + 1
    WHERE site_id = ?
    AND location_country_name = ?
    AND location_country_code = ?
    AND location_country_id = ?
    AND made_at_year = ?
    AND made_at_month = ?
    AND made_at_day = ?
    `
	return sdcacr.Repository.
		Cassandra.
		Query(cql, siteID, countryName, countryCode, countryID, date.Year(), date.Month(), date.Day()).
		Exec()
}
