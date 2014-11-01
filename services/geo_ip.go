package services

import (
	geoip2 "github.com/oschwald/geoip2-golang"
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
)

// Singleton instance of GeoIP database.
var GeoIP *geoip2.Reader

// InitGeoIP opens GeoIP database file.
func InitGeoIP(config lib.GeoIPConfig) {
	geoIP, err := geoip2.Open(config.Path)

	if err != nil {
		Logger.Error("GeoIP database cannot be opened.")
		panic(err)
	}

	Logger.Info("GeoIP database was opened successfully.")

	GeoIP = geoIP
}
