package repositories

import (
	"time"

	"github.com/gocql/gocql"
)

const (
	// MetricDayByMinuteCounterFields ...
	MetricDayByMinuteCounterFields = `
    dimensions_names, dimensions_values, made_at_minute, made_at_day,
	made_at_month, made_at_year, count,
	`
)

// MetricDayByMinuteCounterRepository ...
type MetricDayByMinuteCounterRepository struct {
	Repository
	ColumnFamily string
}

// NewMetricDayByMinuteCounterRepository ...
func NewMetricDayByMinuteCounterRepository(repository Repository, columnFamily string) MetricDayByMinuteCounterRepository {
	return MetricDayByMinuteCounterRepository{
		Repository:   repository,
		ColumnFamily: columnFamily,
	}
}

// IncrementQuery ...
func (mdbmcr *MetricDayByMinuteCounterRepository) IncrementQuery(dimensionsNames, dimensionsValues string, now time.Time) *gocql.Query {
	cql := `
    UPDATE ` + mdbmcr.ColumnFamily + `
    SET count = count + 1
	WHERE dimensions_names = ?
    AND dimensions_values = ?
    AND made_at_year = ?
    AND made_at_month = ?
    AND made_at_day = ?
    AND made_at_minute = ?
    `

	from := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	nbOfMinutes := int(now.Sub(from).Minutes())

	return mdbmcr.Repository.
		Cassandra.
		Query(cql, dimensionsNames, dimensionsValues, now.Year(), now.Month(), now.Day(), nbOfMinutes)
}

// Increment ...
func (mdbmcr *MetricDayByMinuteCounterRepository) Increment(dimensionsNames, dimensionsValues string, date time.Time) error {
	return mdbmcr.IncrementQuery(dimensionsNames, dimensionsValues, date).Exec()
}
