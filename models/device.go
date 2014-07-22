package models

type Device struct {
	Name     string `json:"name" bson:"name"`
	IsMobile bool   `json:"isMobile" bson:"is_mobile"`
	IsTablet bool   `json:"isTablet" bson:"is_tablet"`
	IsPhone  bool   `json:"isPhone" bson:"is_phone"`
}
