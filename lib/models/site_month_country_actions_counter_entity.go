package models

// SiteMonthCountryActionsCounterEntity ...
type SiteMonthCountryActionsCounterEntity struct {
	SiteID              int64  `cql:"site_id"`
	LocationCountryName string `cql:"location_country_name"`
	LocationCountryCode string `cql:"location_country_code"`
	LocationCountryID   uint   `cql:"location_country_id"`
	MadeAtYear          int    `cql:"made_at_year"`
	MadeAtMonth         int    `cql:"made_at_month"`
	NbOfActions         int64  `cql:"nb_of_actions"`
}
