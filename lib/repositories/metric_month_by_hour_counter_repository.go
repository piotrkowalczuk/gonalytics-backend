package repositories

import (
	"time"

	"github.com/gocql/gocql"
)

const (
	// MetricMonthByHourCounterFields ...
	MetricMonthByHourCounterFields = `
    dimensions_names, dimensions_values, made_at_hour,
	made_at_month, made_at_year, count,
	`
)

// MetricMonthByHourCounterRepository ...
type MetricMonthByHourCounterRepository struct {
	Repository
	ColumnFamily string
}

// NewMetricMonthByHourCounterRepository ...
func NewMetricMonthByHourCounterRepository(repository Repository, columnFamily string) MetricMonthByHourCounterRepository {
	return MetricMonthByHourCounterRepository{
		Repository:   repository,
		ColumnFamily: columnFamily,
	}
}

// IncrementQuery ...
func (mmbhcr *MetricMonthByHourCounterRepository) IncrementQuery(dimensionsNames, dimensionsValues string, now time.Time) *gocql.Query {
	cql := `
    UPDATE ` + mmbhcr.ColumnFamily + `
    SET count = count + 1
	WHERE dimensions_names = ?
    AND dimensions_values = ?
    AND made_at_year = ?
    AND made_at_month = ?
    AND made_at_hour = ?
    `
	from := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	nbOfHours := int(now.Sub(from).Hours())

	return mmbhcr.Repository.
		Cassandra.
		Query(cql, dimensionsNames, dimensionsValues, now.Year(), now.Month(), nbOfHours)
}

// Increment ...
func (mmbhcr *MetricMonthByHourCounterRepository) Increment(dimensionsNames, dimensionsValues string, date time.Time) error {
	return mmbhcr.IncrementQuery(dimensionsNames, dimensionsValues, date).Exec()
}
