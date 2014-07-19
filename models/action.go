package models

import (
	"github.com/oschwald/geoip2-golang"
)

type Action struct {
	Id              string           `json:"id" bson:"_id,omitempty"`
	UserId          string           `json:"userId" bson:"userId"`
	SiteId          int64            `json:"siteId" bson:"siteId"`
	Referrer        string           `json:"referrer" bson:"referrer"`
	Language        string           `json:"language" bson:"language"`
	Browser         *Browser         `json:"browser" bson:"browser"`
	Screen          *Screen          `json:"screen" bson:"screen"`
	Website         *Website         `json:"website" bson:"website"`
	OperatingSystem *OperatingSystem `json:"os" bson:"os"`
	Device          *Device          `json:"device" bson:"device"`
	Location        *Location        `json:"location" bson:"location"`
	CreatedAt       *MongoDate       `json:"createdAt" bson:"created_at"`
}

type Website struct {
	Title string `json:"title" bson:"title"`
	Host  string `json:"host" bson:"host"`
	Url   string `json:"url" bson:"url"`
}

type Screen struct {
	Width  int64 `json:"width" bson:"width"`
	Height int64 `json:"height" bson:"height"`
}

type Window struct {
	Width  int64 `json:"width" bson:"width"`
	Height int64 `json:"height" bson:"height"`
}

type Browser struct {
	Name         string  `json:"name" bson:"name"`
	Version      string  `json:"version" bson:"version"`
	MajorVersion string  `json:"majorVersion" bson:"major_version"`
	UserAgent    string  `json:"userAgent" bson:"user_agent"`
	Platform     string  `json:"platform" bson:"platform"`
	Cookie       bool    `json:"cookie" bson:"cookie"`
	Plugins      Plugins `json:"plugins" bson:"plugins"`
	IsOnline     bool    `json:"isOnline" bson:"is_online"`
	Window       Window  `json:"window" bson:"window"`
}

type Plugins struct {
	Java bool `json:"java" bson:"java"`
}

type OperatingSystem struct {
	Name    string `json:"name" bson:"name"`
	Version string `json:"version" bson:"version"`
}

type Device struct {
	Name     string `json:"name" bson:"name"`
	IsMobile bool   `json:"isMobile" bson:"is_mobile"`
	IsTablet bool   `json:"isTablet" bson:"is_tablet"`
	IsPhone  bool   `json:"isPhone" bson:"is_phone"`
}

type Location struct {
	CityName            string  `json:"cityName" bson:"city_name"`
	CityId              uint    `json:"cityId" bson:"city_id"`
	CountryName         string  `json:"countryName" bson:"country_name"`
	CountryCode         string  `json:"countryCode" bson:"country_code"`
	CountryId           uint    `json:"countryId" bson:"country_id"`
	ContinentName       string  `json:"continentName" bson:"continent_name"`
	ContinentCode       string  `json:"continentCode" bson:"continent_code"`
	ContinentId         uint    `json:"continentId" bson:"continent_id"`
	Latitude            float64 `json:"latitude" bson:"latitude"`
	Longitude           float64 `json:"longitude" bson:"longitude"`
	MetroCode           uint    `json:"metroCode" bson:"metro_code"`
	TimeZone            string  `json:"timeZone" bson:"time_zone"`
	PostalCode          string  `json:"postalCode" bson:"postal_code"`
	IsAnonymousProxy    bool    `json:"isAnonymousProxy" bson:"is_anonymous_proxy"`
	IsSatelliteProvider bool    `json:"isSatelliteProvider" bson:"is_satellite_provider"`
}

func NewLocationFromGeoIP(geoLocation *geoip2.City) *Location {
	return &Location{
		CityName:            geoLocation.City.Names["en"],
		CityId:              geoLocation.City.GeoNameID,
		CountryName:         geoLocation.Country.Names["en"],
		CountryCode:         geoLocation.Country.IsoCode,
		CountryId:           geoLocation.Country.GeoNameID,
		ContinentName:       geoLocation.Continent.Names["en"],
		ContinentCode:       geoLocation.Continent.Code,
		ContinentId:         geoLocation.Continent.GeoNameID,
		Latitude:            geoLocation.Location.Latitude,
		Longitude:           geoLocation.Location.Longitude,
		MetroCode:           geoLocation.Location.MetroCode,
		TimeZone:            geoLocation.Location.TimeZone,
		PostalCode:          geoLocation.Postal.Code,
		IsAnonymousProxy:    geoLocation.Traits.IsAnonymousProxy,
		IsSatelliteProvider: geoLocation.Traits.IsSatelliteProvider,
	}
}
