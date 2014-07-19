package services

import (
	"github.com/oschwald/geoip2-golang"
	"net"
)

type GeoLocation struct {
	IP       string
	Location *geoip2.City
}

func NewGeoLocation(IP string) (*GeoLocation, error) {
	gl := &GeoLocation{IP: IP}

	db, err := geoip2.Open("GeoLite2-City.mmdb")
	defer db.Close()

	if err != nil {
		return nil, err
	}

	record, err := db.City(net.ParseIP(gl.IP))

	if err != nil {
		return nil, err
	}

	gl.Location = record

	return gl, nil
}
