package services

import geoip2 "github.com/oschwald/geoip2-golang"

// Singleton instance of GeoIP database.
var GeoIP *geoip2.Reader

// InitGeoIP opens GeoIP database file.
func InitGeoIP(filePath string) *geoip2.Reader {
	geoIP, err := geoip2.Open(filePath)

	if err != nil {
		panic(err)
	}

	GeoIP = geoIP
	return geoIP
}
