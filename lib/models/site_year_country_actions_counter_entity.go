package models

// SiteYearCountryActionsCounterEntity ...
type SiteYearCountryActionsCounterEntity struct {
	SiteID              int64  `cql:"site_id"`
	LocationCountryName string `cql:"location_country_name"`
	LocationCountryCode string `cql:"location_country_code"`
	LocationCountryID   uint   `cql:"location_country_id"`
	MadeAtYear          int    `cql:"made_at_year"`
	NbOfActions         int64  `cql:"nb_of_actions"`
}
