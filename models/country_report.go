package models

type CountryReport struct {
	NbOfReports int64                      `json:"nbOfReports" xml:"nbOfReports"`
	Reports     map[string]*CountrySummary `json:"reports" xml:"reports"`
}

type CountrySummary struct {
	CountryName        string `json:"countryName" xml:"countryName"`
	CountryCode        string `json:"countryCode" xml:"countryCode"`
	CountryId          uint   `json:"countryId" xml:"countryId"`
	NbOfActions        int64  `json:"nbOfActions" xml:"nbOfActions"`
	NbOfVisits         int64  `json:"nbOfVisits" xml:"nbOfVisits"`
	NbOfUniqueVisitors int64  `json:"nbOfUniqueVisitors" xml:"nbOfUniqueVisitors"`
	SiteId             int64  `json:"siteId" xml:"siteId"`
}

func NewCountryReportFromVisits(visits []*Visit) *CountryReport {
	countryReport := CountryReport{
		NbOfReports: 0,
		Reports:     make(map[string]*CountrySummary),
	}

	for _, visit := range visits {
		countryCode := visit.Location.CountryCode

		if _, exists := countryReport.Reports[countryCode]; exists {
			countryReport.Reports[countryCode].NbOfActions += visit.NbOfActions
			countryReport.Reports[countryCode].NbOfVisits++
			countryReport.Reports[countryCode].NbOfUniqueVisitors++
		} else {
			countryReport.Reports[countryCode] = &CountrySummary{
				CountryName:        visit.Location.CountryName,
				CountryCode:        visit.Location.CountryCode,
				CountryId:          visit.Location.CountryId,
				NbOfActions:        visit.NbOfActions,
				NbOfVisits:         1,
				NbOfUniqueVisitors: 1,
				SiteId:             visit.SiteId,
			}
		}
	}

	return &countryReport
}
