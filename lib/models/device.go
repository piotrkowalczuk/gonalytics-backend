package models

// Device ...
type Device struct {
	Name     string `json:"name" cql:"name"`
	IsMobile bool   `json:"isMobile" cql:"is_mobile"`
	IsTablet bool   `json:"isTablet" cql:"is_tablet"`
	IsPhone  bool   `json:"isPhone" cql:"is_phone"`
}
