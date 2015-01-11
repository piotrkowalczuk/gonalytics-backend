package models

// SiteMonthCountryActionsCounterEntity ...
type SiteMonthCountryActionsCounterEntity struct {
	SiteID              int64  `json:"siteId" cql:"site_id"`
	LocationCountryName string `json:"locationCountryName" cql:"location_country_name"`
	LocationCountryCode string `json:"locationCountryCode" cql:"location_country_code"`
	LocationCountryID   uint   `json:"locationCountryId" cql:"location_country_id"`
	MadeAtYear          int    `json:"madeAtYear" cql:"made_at_year"`
	MadeAtMonth         int    `json:"madeAtMonth" cql:"made_at_month"`
	NbOfActions         int64  `json:"nbOfActions" cql:"nb_of_actions"`
}
