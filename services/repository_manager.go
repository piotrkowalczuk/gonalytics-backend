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
			Repository: repository,
		},
		// Actions by browser
		SiteDayBrowserActionsCounter: repositories.SiteDayBrowserCounterRepository{
			Repository: repositories.Repository{
				Cassandra:    cassandra,
				ColumnFamily: repositories.SiteDayBrowserActionsCounterColumnFamily,
			},
		},
		SiteMonthBrowserActionsCounter: repositories.SiteMonthBrowserCounterRepository{
			Repository: repositories.Repository{
				Cassandra:    cassandra,
				ColumnFamily: repositories.SiteMonthBrowserActionsCounterColumnFamily,
			},
		},
		SiteYearBrowserActionsCounter: repositories.SiteYearBrowserCounterRepository{
			Repository: repositories.Repository{
				Cassandra:    cassandra,
				ColumnFamily: repositories.SiteYearBrowserActionsCounterColumnFamily,
			},
		},
		// Visits by browser
		SiteDayBrowserVisitsCounter: repositories.SiteDayBrowserCounterRepository{
			Repository: repositories.Repository{
				Cassandra:    cassandra,
				ColumnFamily: repositories.SiteDayBrowserVisitsCounterColumnFamily,
			},
		},
		SiteMonthBrowserVisitsCounter: repositories.SiteMonthBrowserCounterRepository{
			Repository: repositories.Repository{
				Cassandra:    cassandra,
				ColumnFamily: repositories.SiteMonthBrowserVisitsCounterColumnFamily,
			},
		},
		SiteYearBrowserVisitsCounter: repositories.SiteYearBrowserCounterRepository{
			Repository: repositories.Repository{
				Cassandra:    cassandra,
				ColumnFamily: repositories.SiteYearBrowserVisitsCounterColumnFamily,
			},
		},
		// Actions by country
		SiteDayCountryActionsCounter: repositories.SiteDayCountryCounterRepository{
			Repository: repositories.Repository{
				Cassandra:    cassandra,
				ColumnFamily: repositories.SiteDayCountryActionsCounterColumnFamily,
			},
		},
		SiteMonthCountryActionsCounter: repositories.SiteMonthCountryCounterRepository{
			Repository: repositories.Repository{
				Cassandra:    cassandra,
				ColumnFamily: repositories.SiteMonthCountryActionsCounterColumnFamily,
			},
		},
		SiteYearCountryActionsCounter: repositories.SiteYearCountryCounterRepository{
			Repository: repositories.Repository{
				Cassandra:    cassandra,
				ColumnFamily: repositories.SiteYearCountryActionsCounterColumnFamily,
			},
		},
		// Visits by country
		SiteDayCountryVisitsCounter: repositories.SiteDayCountryCounterRepository{
			Repository: repositories.Repository{
				Cassandra:    cassandra,
				ColumnFamily: repositories.SiteDayCountryVisitsCounterColumnFamily,
			},
		},
		SiteMonthCountryVisitsCounter: repositories.SiteMonthCountryCounterRepository{
			Repository: repositories.Repository{
				Cassandra:    cassandra,
				ColumnFamily: repositories.SiteMonthCountryVisitsCounterColumnFamily,
			},
		},
		SiteYearCountryVisitsCounter: repositories.SiteYearCountryCounterRepository{
			Repository: repositories.Repository{
				Cassandra:    cassandra,
				ColumnFamily: repositories.SiteYearCountryVisitsCounterColumnFamily,
			},
		},
	}
}
