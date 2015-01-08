package repositories

import "time"

const (
	// SiteYearCountryActionsCounterColumnFamily ...
	SiteYearCountryActionsCounterColumnFamily = "site_year_country_actions_counter"
)

// SiteYearCountryActionsCounterRepository ...
type SiteYearCountryActionsCounterRepository struct {
	Repository
}

// Increment ...
func (sycacr *SiteYearCountryActionsCounterRepository) Increment(siteID int64, countryName string, countryCode string, countryID uint, date time.Time) error {

	cql := `
    UPDATE ` + SiteYearCountryActionsCounterColumnFamily + `
    SET nb_of_actions = nb_of_actions + 1
    WHERE site_id = ?
    AND location_country_name = ?
    AND location_country_code = ?
    AND location_country_id = ?
    AND made_at_year = ?
    `
	return sycacr.Repository.
		Cassandra.
		Query(cql, siteID, countryName, countryCode, countryID, date.Year()).
		Exec()
}
