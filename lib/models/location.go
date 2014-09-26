package models

import (
	geoip2 "github.com/oschwald/geoip2-golang"
)

// Location ...
type Location struct {
	CityName            string  `json:"cityName" bson:"city_name"`
	CityID              uint    `json:"cityId" bson:"city_id"`
	CountryName         string  `json:"countryName" bson:"country_name"`
	CountryCode         string  `json:"countryCode" bson:"country_code"`
	CountryID           uint    `json:"countryId" bson:"country_id"`
	ContinentName       string  `json:"continentName" bson:"continent_name"`
	ContinentCode       string  `json:"continentCode" bson:"continent_code"`
	ContinentID         uint    `json:"continentId" bson:"continent_id"`
	Latitude            float64 `json:"latitude" bson:"latitude"`
	Longitude           float64 `json:"longitude" bson:"longitude"`
	MetroCode           uint    `json:"metroCode" bson:"metro_code"`
	TimeZone            string  `json:"timeZone" bson:"time_zone"`
	PostalCode          string  `json:"postalCode" bson:"postal_code"`
	IsAnonymousProxy    bool    `json:"isAnonymousProxy" bson:"is_anonymous_proxy"`
	IsSatelliteProvider bool    `json:"isSatelliteProvider" bson:"is_satellite_provider"`
}

// NewLocationFromGeoIP ...
func NewLocationFromGeoIP(geoLocation *geoip2.City) *Location {
	return &Location{
		CityName:            geoLocation.City.Names["en"],
		CityID:              geoLocation.City.GeoNameID,
		CountryName:         geoLocation.Country.Names["en"],
		CountryCode:         geoLocation.Country.IsoCode,
		CountryID:           geoLocation.Country.GeoNameID,
		ContinentName:       geoLocation.Continent.Names["en"],
		ContinentCode:       geoLocation.Continent.Code,
		ContinentID:         geoLocation.Continent.GeoNameID,
		Latitude:            geoLocation.Location.Latitude,
		Longitude:           geoLocation.Location.Longitude,
		MetroCode:           geoLocation.Location.MetroCode,
		TimeZone:            geoLocation.Location.TimeZone,
		PostalCode:          geoLocation.Postal.Code,
		IsAnonymousProxy:    geoLocation.Traits.IsAnonymousProxy,
		IsSatelliteProvider: geoLocation.Traits.IsSatelliteProvider,
	}
}
