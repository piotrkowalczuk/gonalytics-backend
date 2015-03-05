package lib

import (
	"net"
	"time"

	"github.com/gocql/gocql"
	geoip2 "github.com/oschwald/geoip2-golang"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
)

// ActionCreator ...
type ActionCreator struct {
	geoIP        *geoip2.Reader
	trackRequest *TrackRequest
}

// NewActionCreator ...
func NewActionCreator(geoIP *geoip2.Reader) *ActionCreator {
	return &ActionCreator{
		geoIP: geoIP,
	}
}

// Create creates visit object based on given track request.
func (ac *ActionCreator) Create(trackRequest *TrackRequest) (*models.ActionEntity, error) {
	ac.trackRequest = trackRequest
	location, err := ac.createLocation()

	if err != nil {
		return nil, err
	}

	var visitID gocql.UUID
	visitID, err = gocql.ParseUUID(ac.trackRequest.VisitID)
	if err != nil {
		return nil, err
	}

	timeLocation, err := time.LoadLocation(location.Timezone)
	if err != nil {
		return nil, err
	}

	locationTime := ac.trackRequest.ServerTime.In(timeLocation)

	return &models.ActionEntity{
		ID:                          gocql.TimeUUID(),
		VisitID:                     visitID,
		IP:                          ac.trackRequest.GetRequestIP(),
		Referrer:                    ac.trackRequest.Referrer,
		Language:                    ac.trackRequest.Language,
		SiteID:                      ac.trackRequest.SiteID,
		LocationCityName:            location.CityName,
		LocationCityID:              location.CityID,
		LocationCountryName:         location.CountryName,
		LocationCountryCode:         location.CountryCode,
		LocationCountryID:           location.CountryID,
		LocationContinentName:       location.ContinentName,
		LocationContinentCode:       location.ContinentCode,
		LocationContinentID:         location.ContinentID,
		LocationLatitude:            location.Latitude,
		LocationLongitude:           location.Longitude,
		LocationMetroCode:           location.MetroCode,
		LocationTime:                locationTime,
		LocationTimezone:            location.Timezone,
		LocationPostalCode:          location.PostalCode,
		LocationIsAnonymousProxy:    location.IsAnonymousProxy,
		LocationIsSatelliteProvider: location.IsSatelliteProvider,
		BrowserName:                 ac.trackRequest.BrowserName,
		BrowserVersion:              ac.trackRequest.BrowserVersion,
		BrowserMajorVersion:         ac.trackRequest.BrowserMajorVersion,
		BrowserUserAgent:            ac.trackRequest.BrowserUserAgent,
		BrowserPlatform:             ac.trackRequest.BrowserPlatform,
		BrowserCookie:               ac.trackRequest.BrowserCookie,
		BrowserIsOnline:             ac.trackRequest.BrowserIsOnline,
		BrowserPluginJava:           ac.trackRequest.BrowserPluginJava,
		BrowserWindowWidth:          ac.trackRequest.BrowserWindowWidth,
		BrowserWindowHeight:         ac.trackRequest.BrowserWindowHeight,
		PageTitle:                   ac.trackRequest.PageTitle,
		PageHost:                    ac.trackRequest.PageHost,
		PageURL:                     ac.trackRequest.PageURL,
		OperatingSystemName:         ac.trackRequest.OperatingSystemName,
		OperatingSystemVersion:      ac.trackRequest.OperatingSystemVersion,
		ScreenWidth:                 ac.trackRequest.ScreenWidth,
		ScreenHeight:                ac.trackRequest.ScreenHeight,
		DeviceName:                  ac.trackRequest.DeviceName,
		DeviceIsTablet:              ac.trackRequest.DeviceIsTablet,
		DeviceIsPhone:               ac.trackRequest.DeviceIsPhone,
		DeviceIsMobile:              ac.trackRequest.DeviceIsMobile,
		ServerTime:                  ac.trackRequest.ServerTime,
	}, nil
}

func (ac *ActionCreator) createLocation() (*models.Location, error) {
	// geoLocation, err := ac.geoIP.City(net.ParseIP(ac.trackRequest.GetRequestIP()))
	geoLocation, err := ac.geoIP.City(net.ParseIP("92.225.124.196"))
	if err != nil {
		return nil, err
	}

	return models.NewLocationFromGeoIP(geoLocation), nil
}
