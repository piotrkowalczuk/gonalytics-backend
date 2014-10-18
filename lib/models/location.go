package models

import (
	geoip2 "github.com/oschwald/geoip2-golang"
)

// Location ...
type Location struct {
	CityName            string  `json:"cityName" cql:"city_name"`
	CityID              uint    `json:"cityId" cql:"city_id"`
	CountryName         string  `json:"countryName" cql:"country_name"`
	CountryCode         string  `json:"countryCode" cql:"country_code"`
	CountryID           uint    `json:"countryId" cql:"country_id"`
	ContinentName       string  `json:"continentName" cql:"continent_name"`
	ContinentCode       string  `json:"continentCode" cql:"continent_code"`
	ContinentID         uint    `json:"continentId" cql:"continent_id"`
	Latitude            float64 `json:"latitude" cql:"latitude"`
	Longitude           float64 `json:"longitude" cql:"longitude"`
	MetroCode           uint    `json:"metroCode" cql:"metro_code"`
	TimeZone            string  `json:"timeZone" cql:"time_zone"`
	PostalCode          string  `json:"postalCode" cql:"postal_code"`
	IsAnonymousProxy    bool    `json:"isAnonymousProxy" cql:"is_anonymous_proxy"`
	IsSatelliteProvider bool    `json:"isSatelliteProvider" cql:"is_satellite_provider"`
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
