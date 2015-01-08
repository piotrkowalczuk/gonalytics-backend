package lib

import "github.com/piotrkowalczuk/gonalytics-backend/lib/repositories"

// RepositoryManager ...
type RepositoryManager struct {
	VisitAction                    repositories.VisitActionRepository
	SiteDayCountryActionsCounter   repositories.SiteDayCountryActionsCounterRepository
	SiteMonthCountryActionsCounter repositories.SiteMonthCountryActionsCounterRepository
	SiteYearCountryActionsCounter  repositories.SiteYearCountryActionsCounterRepository
}
