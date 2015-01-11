package models

// SiteDayCountryCounter ...
type SiteDayCountryCounter struct {
	SiteID              int64  `json:"siteId"`
	LocationCountryName string `json:"locationCountryName"`
	LocationCountryCode string `json:"locationCountryCode"`
	LocationCountryID   string `json:"locationCountryID"`
	MadeAtYear          int    `json:"madeAtYear"`
	MadeAtMonth         int    `json:"madeAtMonth"`
	MadeAtDay           int    `json:"madeAtDay"`
	Count               int64  `json:"count"`
}

// SiteDayCountryActionsCounterEntity ...
type SiteDayCountryActionsCounterEntity struct {
	SiteDayCountryCounter
}

// SiteDayCountryVisitsCounterEntity ...
type SiteDayCountryVisitsCounterEntity struct {
	SiteDayCountryCounter
}

// SiteMonthCountryCounter ...
type SiteMonthCountryCounter struct {
	SiteID              int64  `json:"siteId"`
	LocationCountryName string `json:"locationCountryName"`
	LocationCountryCode string `json:"locationCountryCode"`
	LocationCountryID   string `json:"locationCountryID"`
	MadeAtYear          int    `json:"madeAtYear"`
	MadeAtMonth         int    `json:"madeAtMonth"`
	Count               int64  `json:"count"`
}

// SiteMonthCountryActionsCounterEntity ...
type SiteMonthCountryActionsCounterEntity struct {
	SiteMonthCountryCounter
}

// SiteMonthCountryVisitsCounterEntity ...
type SiteMonthCountryVisitsCounterEntity struct {
	SiteMonthCountryCounter
}

// SiteYearCountryCounter ...
type SiteYearCountryCounter struct {
	SiteID              int64  `json:"siteId"`
	LocationCountryName string `json:"locationCountryName"`
	LocationCountryCode string `json:"locationCountryCode"`
	LocationCountryID   string `json:"locationCountryID"`
	MadeAtYear          int    `json:"madeAtYear"`
	Count               int64  `json:"count"`
}

// SiteSiteYearCountryCounterEntityCountryActionsCounterEntity ...
type SiteYearCountryActionsCounterEntity struct {
	SiteYearCountryCounter
}

// SiteYearCountryVisitsCounterEntity ...
type SiteYearCountryVisitsCounterEntity struct {
	SiteYearCountryCounter
}
