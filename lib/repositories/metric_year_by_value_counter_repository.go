package repositories

import (
	"time"

	"github.com/gocql/gocql"
)

const (
	// MetricYearByValueCounterFields ...
	MetricYearByValueCounterFields = `
    dimensions_names, dimensions_values, made_at_year, count,
	`
)

// MetricYearByValueCounterRepository ...
type MetricYearByValueCounterRepository struct {
	Repository
	ColumnFamily string
}

// NewMetricYearByValueCounterRepository ...
func NewMetricYearByValueCounterRepository(repository Repository, columnFamily string) MetricYearByValueCounterRepository {
	return MetricYearByValueCounterRepository{
		Repository:   repository,
		ColumnFamily: columnFamily,
	}
}

// IncrementQuery ...
func (mybvcr *MetricYearByValueCounterRepository) IncrementQuery(dimensionsNames, dimensionsValues string, date time.Time) *gocql.Query {
	cql := `
    UPDATE ` + mybvcr.ColumnFamily + `
    SET count = count + 1
	WHERE dimensions_names = ?
    AND dimensions_values = ?
    AND made_at_year = ?
    `

	return mybvcr.Repository.
		Cassandra.
		Query(cql, dimensionsNames, dimensionsValues, date.Year())
}

// Increment ...
func (mybvcr *MetricYearByValueCounterRepository) Increment(dimensionsNames, dimensionsValues string, date time.Time) error {
	return mybvcr.IncrementQuery(dimensionsNames, dimensionsValues, date).Exec()
}
