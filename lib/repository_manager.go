package lib

import "github.com/piotrkowalczuk/gonalytics-backend/lib/repositories"

// RepositoryManager ...
type RepositoryManager struct {
	VisitActions repositories.VisitActionsRepository
	// Metric by value
	MetricDayByValueCounter   repositories.MetricDayByValueCounterRepository
	MetricMonthByValueCounter repositories.MetricMonthByValueCounterRepository
	MetricYearByValueCounter  repositories.MetricYearByValueCounterRepository
	// Metric by time frame
	MetricDayByMinuteCounter repositories.MetricDayByMinuteCounterRepository
	MetricMonthByHourCounter repositories.MetricMonthByHourCounterRepository
	MetricYearByDayCounter   repositories.MetricYearByDayCounterRepository
}
