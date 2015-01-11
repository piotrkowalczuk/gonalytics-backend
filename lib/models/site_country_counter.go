package models

// SiteYearCountryCounterEnity ...
type SiteYearCountryCounterEntity struct {
	SiteID              int64  `json:"siteId"`
	LocationCountryName string `json:"locationCountryName"`
	LocationCountryCode string `json:"locationCountryCode"`
	LocationCountryID   string `json:"locationCountryID"`
	MadeAtYear          int    `json:"madeAtYear"`
	Count               int64  `json:"count"`
}

// SiteMonthCountryCounterEntity ...
type SiteMonthCountryCounterEntity struct {
	SiteYearCountryCounterEntity
	MadeAtMonth int `json:"madeAtMonth"`
}

// SiteDayCountryCounterEntity ...
type SiteDayCountryCounterEntity struct {
	SiteYearCountryCounterEntity
	MadeAtMonth int `json:"madeAtMonth"`
	MadeAtDay   int `json:"madeAtDay"`
}
