package repositories

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
)

const (
	// SiteMonthCountryActionsCounterColumnFamily ...
	SiteMonthCountryActionsCounterColumnFamily = "site_month_country_actions_counter"
	// SiteMonthCountryActionsFields ...
	SiteMonthCountryActionsFields = `
        site_id, nb_of_actions, location_country_name, location_country_code, location_country_id,
        made_at_year, made_at_month
    `
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

// Find ...
func (smcacr *SiteMonthCountryActionsCounterRepository) Find(
	siteID int64,
	date time.Time,
) ([]*models.SiteMonthCountryActionsCounterEntity, error) {
	cql := `SELECT ` + SiteMonthCountryActionsFields +
		` FROM ` + SiteMonthCountryActionsCounterColumnFamily +
		` WHERE site_id = ? AND made_at_year = ? AND made_at_month = ?`

	iter := smcacr.Repository.
		Cassandra.
		Query(cql, siteID, date.Year(), date.Month()).
		Consistency(gocql.One).
		Iter()

	var counter models.SiteMonthCountryActionsCounterEntity
	counters := []*models.SiteMonthCountryActionsCounterEntity{}

	for iter.Scan(
		&counter.SiteID,
		&counter.NbOfActions,
		&counter.LocationCountryName,
		&counter.LocationCountryCode,
		&counter.LocationCountryID,
		&counter.MadeAtYear,
		&counter.MadeAtMonth,
	) {
		counters = append(counters, &counter)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return counters, nil
}
