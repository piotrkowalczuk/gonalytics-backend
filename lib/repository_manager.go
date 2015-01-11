package lib

import "github.com/piotrkowalczuk/gonalytics-backend/lib/repositories"

// RepositoryManager ...
type RepositoryManager struct {
	VisitActions repositories.VisitActionsRepository
	// Browser
	SiteDayBrowserActionsCounter   repositories.SiteDayBrowserActionsCounterRepository
	SiteMonthBrowserActionsCounter repositories.SiteMonthBrowserActionsCounterRepository
	SiteYearBrowserActionsCounter  repositories.SiteYearBrowserActionsCounterRepository
	// Country
	SiteDayCountryActionsCounter   repositories.SiteDayCountryActionsCounterRepository
	SiteMonthCountryActionsCounter repositories.SiteMonthCountryActionsCounterRepository
	SiteYearCountryActionsCounter  repositories.SiteYearCountryActionsCounterRepository
}
