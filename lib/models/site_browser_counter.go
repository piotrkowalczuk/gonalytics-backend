package models

// SiteDayBrowserCounter ...
type SiteDayBrowserCounter struct {
	SiteID         int64  `json:"siteId"`
	BrowserName    string `json:"browserName"`
	BrowserVersion string `json:"browserVersion"`
	MadeAtYear     int    `json:"madeAtYear"`
	MadeAtMonth    int    `json:"madeAtMonth"`
	MadeAtDay      int    `json:"madeAtDay"`
	Count          int64  `json:"count"`
}

// SiteDayBrowserActionsCounterEntity ...
type SiteDayBrowserActionsCounterEntity struct {
	SiteDayBrowserCounter
}

// SiteDayBrowserVisitsCounterEntity ...
type SiteDayBrowserVisitsCounterEntity struct {
	SiteDayBrowserCounter
}

// SiteMonthBrowserCounter ...
type SiteMonthBrowserCounter struct {
	SiteID         int64  `json:"siteId"`
	BrowserName    string `json:"browserName"`
	BrowserVersion string `json:"browserVersion"`
	MadeAtYear     int    `json:"madeAtYear"`
	MadeAtMonth    int    `json:"madeAtMonth"`
	Count          int64  `json:"count"`
}

// SiteMonthBrowserActionsCounterEntity ...
type SiteMonthBrowserActionsCounterEntity struct {
	SiteMonthBrowserCounter
}

// SiteMonthBrowserVisitsCounterEntity ...
type SiteMonthBrowserVisitsCounterEntity struct {
	SiteMonthBrowserCounter
}

// SiteYearBrowserCounter ...
type SiteYearBrowserCounter struct {
	SiteID         int64  `json:"siteId"`
	BrowserName    string `json:"browserName"`
	BrowserVersion string `json:"browserVersion"`
	MadeAtYear     int    `json:"madeAtYear"`
	Count          int64  `json:"count"`
}

// SiteSiteYearBrowserCounterEntityBrowserActionsCounterEntity ...
type SiteYearBrowserActionsCounterEntity struct {
	SiteYearBrowserCounter
}

// SiteYearBrowserVisitsCounterEntity ...
type SiteYearBrowserVisitsCounterEntity struct {
	SiteYearBrowserCounter
}
