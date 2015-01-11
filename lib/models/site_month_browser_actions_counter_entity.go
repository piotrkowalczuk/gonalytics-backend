package models

// SiteMonthBrowserActionsCounterEntity ...
type SiteMonthBrowserActionsCounterEntity struct {
	SiteID         int64  `json:"siteId" cql:"site_id"`
	BrowserName    string `json:"browserName" cql:"browser_name"`
	BrowserVersion string `json:"browserVersion" cql:"browser_version"`
	MadeAtYear     int    `json:"madeAtYear" cql:"made_at_year"`
	MadeAtMonth    int    `json:"madeAtMonth" cql:"made_at_month"`
	NbOfActions    int64  `json:"nbOfActions" cql:"nb_of_actions"`
}
