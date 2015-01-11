package repositories

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
)

const (
	// SiteDayCountryActionsCounterColumnFamily ...
	SiteDayCountryActionsCounterColumnFamily = "site_day_country_actions_counter"
	// SiteDayCountryActionsFields ...
	SiteDayCountryActionsFields = `
        site_id, nb_of_actions, location_country_name, location_country_code, location_country_id,
        made_at_year, made_at_month, made_at_day
    `
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

// Find ...
func (sdcacr *SiteDayCountryActionsCounterRepository) Find(
	siteID int64,
	date time.Time,
) ([]*models.SiteDayCountryActionsCounterEntity, error) {
	cql := `SELECT ` + SiteDayCountryActionsFields +
		` FROM ` + SiteDayCountryActionsCounterColumnFamily +
		` WHERE site_id = ? AND made_at_year = ? AND made_at_month = ? AND made_at_day = ?`

	iter := sdcacr.Repository.
		Cassandra.
		Query(cql, siteID, date.Year(), date.Month(), date.Day()).
		Consistency(gocql.One).
		Iter()

	var counter models.SiteDayCountryActionsCounterEntity
	counters := []*models.SiteDayCountryActionsCounterEntity{}

	for iter.Scan(
		&counter.SiteID,
		&counter.NbOfActions,
		&counter.LocationCountryName,
		&counter.LocationCountryCode,
		&counter.LocationCountryID,
		&counter.MadeAtYear,
		&counter.MadeAtMonth,
		&counter.MadeAtDay,
	) {
		counters = append(counters, &counter)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return counters, nil
}
