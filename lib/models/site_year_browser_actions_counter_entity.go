package models

// SiteYearBrowserActionsCounterEntity ...
type SiteYearBrowserActionsCounterEntity struct {
	SiteID         int64  `cql:"site_id"`
	BrowserName    string `cql:"browser_name"`
	BrowserVersion string `cql:"browser_version"`
	MadeAtYear     int    `cql:"made_at_year"`
	NbOfActions    int64  `cql:"nb_of_actions"`
}
