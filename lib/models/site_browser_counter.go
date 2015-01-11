package models

// SiteYearBrowserCounter ...
type SiteYearBrowserCounterEntity struct {
	SiteID         int64  `json:"siteId"`
	BrowserName    string `json:"browserName"`
	BrowserVersion string `json:"browserVersion"`
	MadeAtYear     int    `json:"madeAtYear"`
	Count          int64  `json:"count"`
}

// SiteMonthBrowserCounterEntity ...
type SiteMonthBrowserCounterEntity struct {
	SiteYearBrowserCounterEntity
	MadeAtMonth int `json:"madeAtMonth"`
}

// SiteDayBrowserCounter ...
type SiteDayBrowserCounterEntity struct {
	SiteYearBrowserCounterEntity
	MadeAtMonth int `json:"madeAtMonth"`
	MadeAtDay   int `json:"madeAtDay"`
}
