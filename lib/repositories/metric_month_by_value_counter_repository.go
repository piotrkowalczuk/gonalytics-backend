package repositories

import (
	"time"

	"github.com/gocql/gocql"
)

const (
	// MetricMonthByValueCounterFields ...
	MetricMonthByValueCounterFields = `
    dimensions_names, dimensions_values, made_at_month, made_at_year, count,
	`
)

// MetricMonthByValueCounterRepository ...
type MetricMonthByValueCounterRepository struct {
	Repository
	ColumnFamily string
}

// NewMetricMonthByValueCounterRepository ...
func NewMetricMonthByValueCounterRepository(repository Repository, columnFamily string) MetricMonthByValueCounterRepository {
	return MetricMonthByValueCounterRepository{
		Repository:   repository,
		ColumnFamily: columnFamily,
	}
}

// IncrementQuery ...
func (mmbvcr *MetricMonthByValueCounterRepository) IncrementQuery(dimensionsNames, dimensionsValues string, date time.Time) *gocql.Query {
	cql := `
    UPDATE ` + mmbvcr.ColumnFamily + `
    SET count = count + 1
	WHERE dimensions_names = ?
    AND dimensions_values = ?
    AND made_at_year = ?
    AND made_at_month = ?
    `

	return mmbvcr.Repository.
		Cassandra.
		Query(cql, dimensionsNames, dimensionsValues, date.Year(), date.Month())
}

// Increment ...
func (mmbvcr *MetricMonthByValueCounterRepository) Increment(dimensionsNames, dimensionsValues string, date time.Time) error {
	return mmbvcr.IncrementQuery(dimensionsNames, dimensionsValues, date).Exec()
}
