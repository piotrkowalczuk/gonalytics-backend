package lib

import "github.com/piotrkowalczuk/gonalytics-backend/lib/repositories"

// RepositoryManager ...
type RepositoryManager struct {
	VisitActions             repositories.VisitActionsRepository
	MetricDayByValueCounter  repositories.MetricDayByValueCounterRepository
	MetricDayByMinuteCounter repositories.MetricDayByMinuteCounterRepository
	MetricMonthByHourCounter repositories.MetricMonthByHourCounterRepository
	MetricYearByDayCounter   repositories.MetricYearByDayCounterRepository
}
