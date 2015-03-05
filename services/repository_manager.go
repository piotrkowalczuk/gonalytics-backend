package services

import (
	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/repositories"
)

// Singleton instance of RepositoryManager.
var RepositoryManager lib.RepositoryManager

// InitRepositoryManager ...
func InitRepositoryManager(cassandra *gocql.Session) {
	repository := repositories.Repository{
		Cassandra: cassandra,
	}

	RepositoryManager = lib.RepositoryManager{
		VisitActions: repositories.VisitActionsRepository{
			Repository:   repository,
			ColumnFamily: "visit_actions",
		},
		// Metrics by value
		MetricDayByValueCounter:   repositories.NewMetricDayByValueCounterRepository(repository, "metric_day_by_value_counter"),
		MetricMonthByValueCounter: repositories.NewMetricMonthByValueCounterRepository(repository, "metric_month_by_value_counter"),
		MetricYearByValueCounter:  repositories.NewMetricYearByValueCounterRepository(repository, "metric_year_by_value_counter"),
		// Metrics by time frame
		MetricDayByMinuteCounter: repositories.NewMetricDayByMinuteCounterRepository(repository, "metric_day_by_minute_counter"),
		MetricMonthByHourCounter: repositories.NewMetricMonthByHourCounterRepository(repository, "metric_month_by_hour_counter"),
		MetricYearByDayCounter:   repositories.NewMetricYearByDayCounterRepository(repository, "metric_year_by_day_counter"),
	}
}
