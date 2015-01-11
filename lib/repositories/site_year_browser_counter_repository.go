package repositories

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
)

const (
	// SiteYearBrowserActionCounterColumnFamily ...
	SiteYearBrowserActionsCounterColumnFamily = "site_year_browser_actions_counter"
	// SiteYearBrowserVisitsCounterColumnFamily ...
	SiteYearBrowserVisitsCounterColumnFamily = "site_year_browser_visits_counter"
	// SiteYearBrowserCounterFields ...
	SiteYearBrowserCounterFields = `
        site_id, count, browser_name, browser_version,
        made_at_year
    `
)

// SiteYearBrowserCounterRepository ...
type SiteYearBrowserCounterRepository struct {
	Repository
}

// Increment ...
func (sybcr *SiteYearBrowserCounterRepository) Increment(
	siteID int64,
	browserName string,
	browserVersion string,
	date time.Time,
) error {
	cql := `
    UPDATE ` + sybcr.ColumnFamily + `
    SET count = count + 1
    WHERE site_id = ?
    AND browser_name = ?
    AND browser_version = ?
    AND made_at_year = ?
    `
	return sybcr.Repository.
		Cassandra.
		Query(cql, siteID, browserName, browserVersion, date.Year()).
		Exec()
}

// Find ...
func (sybcr *SiteYearBrowserCounterRepository) Find(
	siteID int64,
	date time.Time,
) ([]*models.SiteYearBrowserCounterEntity, error) {
	cql := `SELECT ` + SiteYearBrowserCounterFields +
		` FROM ` + sybcr.ColumnFamily +
		` WHERE site_id = ? AND made_at_year = ?`

	iter := sybcr.Repository.
		Cassandra.
		Query(cql, siteID, date.Year()).
		Consistency(gocql.One).
		Iter()

	var counter models.SiteYearBrowserCounterEntity
	counters := []*models.SiteYearBrowserCounterEntity{}

	for iter.Scan(
		&counter.SiteID,
		&counter.Count,
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
