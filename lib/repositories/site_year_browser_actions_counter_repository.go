package repositories

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
)

const (
	// SiteYearBrowserActionsCounterColumnFamily ...
	SiteYearBrowserActionsCounterColumnFamily = "site_year_browser_actions_counter"
	// SiteYearBrowserActionsFields ...
	SiteYearBrowserActionsFields = `
        site_id, nb_of_actions, browser_name, browser_version,
        made_at_year
    `
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

// Find ...
func (sybacr *SiteYearBrowserActionsCounterRepository) Find(
	siteID int64,
	date time.Time,
) ([]*models.SiteYearBrowserActionsCounterEntity, error) {
	cql := `SELECT ` + SiteYearBrowserActionsFields +
		` FROM ` + SiteYearBrowserActionsCounterColumnFamily +
		` WHERE site_id = ? AND made_at_year = ?`

	iter := sybacr.Repository.
		Cassandra.
		Query(cql, siteID, date.Year()).
		Consistency(gocql.One).
		Iter()

	var counter models.SiteYearBrowserActionsCounterEntity
	counters := []*models.SiteYearBrowserActionsCounterEntity{}

	for iter.Scan(
		&counter.SiteID,
		&counter.NbOfActions,
		&counter.BrowserName,
		&counter.BrowserVersion,
		&counter.MadeAtYear,
	) {
		counters = append(counters, &counter)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return counters, nil
}
