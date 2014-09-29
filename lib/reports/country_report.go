package reports

import (
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
)

// CountryReports its collections of summaries
type CountryReports struct {
	NbOfreportss int64                      `json:"nbOfreportss" xml:"nbOfreportss"`
	reportss     map[string]*CountrySummary `json:"reportss" xml:"reportss"`
}

// CountrySummary contains basic information about the country
type CountrySummary struct {
	CountryName        string `json:"countryName" xml:"countryName"`
	CountryCode        string `json:"countryCode" xml:"countryCode"`
	CountryID          uint   `json:"countryId" xml:"countryId"`
	NbOfActions        int64  `json:"nbOfActions" xml:"nbOfActions"`
	NbOfVisits         int64  `json:"nbOfVisits" xml:"nbOfVisits"`
	NbOfUniqueVisitors int64  `json:"nbOfUniqueVisitors" xml:"nbOfUniqueVisitors"`
	SiteID             int64  `json:"siteId" xml:"siteId"`
}

// NewCountryReportFromVisits creates reports based on collection of visits
func NewCountryReportFromVisits(visits []*models.Visit) *CountryReports {
	countryreports := CountryReports{
		NbOfreportss: 0,
		reportss:     make(map[string]*CountrySummary),
	}

	for _, visit := range visits {
		countryCode := visit.Location.CountryCode

		if _, exists := countryreports.reportss[countryCode]; exists {
			countryreports.reportss[countryCode].NbOfActions += visit.NbOfActions
			countryreports.reportss[countryCode].NbOfVisits++
			countryreports.reportss[countryCode].NbOfUniqueVisitors++
		} else {
			countryreports.reportss[countryCode] = &CountrySummary{
				CountryName:        visit.Location.CountryName,
				CountryCode:        visit.Location.CountryCode,
				CountryID:          visit.Location.CountryID,
				NbOfActions:        visit.NbOfActions,
				NbOfVisits:         1,
				NbOfUniqueVisitors: 1,
				SiteID:             visit.SiteID,
			}
		}
	}

	return &countryreports
}
