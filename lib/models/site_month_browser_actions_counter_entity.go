package models

// SiteMonthBrowserActionsCounterEntity ...
type SiteMonthBrowserActionsCounterEntity struct {
	SiteID         int64  `cql:"site_id"`
	BrowserName    string `cql:"browser_name"`
	BrowserVersion string `cql:"browser_version"`
	MadeAtYear     int    `cql:"made_at_year"`
	MadeAtMonth    int    `cql:"made_at_month"`
	NbOfActions    int64  `cql:"nb_of_actions"`
}
