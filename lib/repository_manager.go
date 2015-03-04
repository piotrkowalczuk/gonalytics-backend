package lib

import "github.com/piotrkowalczuk/gonalytics-backend/lib/repositories"

// RepositoryManager ...
type RepositoryManager struct {
	VisitActions     repositories.VisitActionsRepository
	MetricDayCounter repositories.MetricDayCounterRepository
}
