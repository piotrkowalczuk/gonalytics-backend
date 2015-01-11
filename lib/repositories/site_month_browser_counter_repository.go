package repositories

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
)

const (
	// SiteMonthBrowserActionsCounterColumnFamily ...
	SiteMonthBrowserActionsCounterColumnFamily = "site_month_browser_actions_counter"
	// SiteMonthBrowserVisitsCounterColumnFamily ...
	SiteMonthBrowserVisitsCounterColumnFamily = "site_month_browser_visits_counter"
	// SiteMonthBrowserActionsFields ...
	SiteMonthBrowserFields = `
        site_id, count, browser_name, browser_version,
        made_at_year, made_at_month
    `
)

// SiteMonthBrowserActionsCounterRepository ...
type SiteMonthBrowserCounterRepository struct {
	Repository
}

// Increment ...
func (smbcr *SiteMonthBrowserCounterRepository) Increment(
	siteID int64,
	browserName string,
	browserVersion string,
	date time.Time,
) error {
	cql := `
    UPDATE ` + smbcr.ColumnFamily + `
    SET count = count + 1
    WHERE site_id = ?
    AND browser_name = ?
    AND browser_version = ?
    AND made_at_year = ?
    AND made_at_month = ?
    `
	return smbcr.Repository.
		Cassandra.
		Query(cql, siteID, browserName, browserVersion, date.Year(), date.Month()).
		Exec()
}

// Find ...
func (smbcr *SiteMonthBrowserCounterRepository) Find(
	siteID int64,
	date time.Time,
) ([]*models.SiteMonthBrowserCounterEntity, error) {
	cql := `SELECT ` + SiteMonthBrowserFields +
		` FROM ` + smbcr.ColumnFamily +
		` WHERE site_id = ? AND made_at_year = ? AND made_at_month = ?`

	iter := smbcr.Repository.
		Cassandra.
		Query(cql, siteID, date.Year(), date.Month()).
		Consistency(gocql.One).
		Iter()

	var counter models.SiteMonthBrowserCounterEntity
	counters := []*models.SiteMonthBrowserCounterEntity{}

	for iter.Scan(
		&counter.SiteID,
		&counter.Count,
		&counter.BrowserName,
		&counter.BrowserVersion,
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
