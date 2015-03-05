package models

import (
	geoip2 "github.com/oschwald/geoip2-golang"
)

// Location ...
type Location struct {
	CityName            string  `json:"cityName"`
	CityID              uint    `json:"cityId"`
	CountryName         string  `json:"countryName"`
	CountryCode         string  `json:"countryCode"`
	CountryID           uint    `json:"countryId"`
	ContinentName       string  `json:"continentName"`
	ContinentCode       string  `json:"continentCode"`
	ContinentID         uint    `json:"continentId"`
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	MetroCode           uint    `json:"metroCode"`
	Timezone            string  `json:"timezone"`
	PostalCode          string  `json:"postalCode"`
	IsAnonymousProxy    bool    `json:"isAnonymousProxy"`
	IsSatelliteProvider bool    `json:"isSatelliteProvider"`
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
		Timezone:            geoLocation.Location.TimeZone,
		PostalCode:          geoLocation.Postal.Code,
		IsAnonymousProxy:    geoLocation.Traits.IsAnonymousProxy,
		IsSatelliteProvider: geoLocation.Traits.IsSatelliteProvider,
	}
}
