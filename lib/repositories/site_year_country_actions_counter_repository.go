package repositories

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
)

const (
	// SiteYearCountryActionsCounterColumnFamily ...
	SiteYearCountryActionsCounterColumnFamily = "site_year_country_actions_counter"
	// SiteYearCountryActionsFields ...
	SiteYearCountryActionsFields = `
        site_id, nb_of_actions, location_country_name, 
        location_country_code, location_country_id,
        made_at_year
    `
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

// Find ...
func (sycacr *SiteYearCountryActionsCounterRepository) Find(
	siteID int64,
	date time.Time,
) ([]*models.SiteYearCountryActionsCounterEntity, error) {
	cql := `SELECT ` + SiteYearCountryActionsFields +
		` FROM ` + SiteYearCountryActionsCounterColumnFamily +
		` WHERE site_id = ? AND made_at_year = ?`

	iter := sycacr.Repository.
		Cassandra.
		Query(cql, siteID, date.Year()).
		Consistency(gocql.One).
		Iter()

	var counter models.SiteYearCountryActionsCounterEntity
	counters := []*models.SiteYearCountryActionsCounterEntity{}

	for iter.Scan(
		&counter.SiteID,
		&counter.NbOfActions,
		&counter.LocationCountryName,
		&counter.LocationCountryCode,
		&counter.LocationCountryID,
		&counter.MadeAtYear,
	) {
		counters = append(counters, &counter)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return counters, nil
}
