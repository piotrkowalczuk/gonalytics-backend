package lib

import (
	"net"

	"github.com/gocql/gocql"
	geoip2 "github.com/oschwald/geoip2-golang"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
)

// VisitCreator ...
type VisitCreator struct {
	geoIP        *geoip2.Reader
	trackRequest *models.TrackRequest
}

// NewVisitCreator ...
func NewVisitCreator(geoIP *geoip2.Reader) *VisitCreator {
	return &VisitCreator{
		geoIP: geoIP,
	}
}

// Create creates visit object based on given track request.
func (vc *VisitCreator) Create(trackRequest *models.TrackRequest) (*models.Visit, error) {
	vc.trackRequest = trackRequest
	location, err := vc.createLocation()

	if err != nil {
		return nil, err
	}

	return &models.Visit{
		ID:              gocql.TimeUUID(),
		IP:              vc.trackRequest.GetRequestIP(),
		Referrer:        vc.trackRequest.Referrer,
		Language:        vc.trackRequest.Language,
		NbOfActions:     1,
		SiteID:          vc.trackRequest.SiteID,
		Location:        location,
		Browser:         vc.createBrowser(),
		FirstPage:       vc.createPage(),
		LastPage:        vc.createPage(),
		OperatingSystem: vc.createOperatingSystem(),
		Screen:          vc.createScreen(),
		Device:          vc.createDevice(),
		FirstActionAt:   vc.trackRequest.MadeAt,
		LastActionAt:    vc.trackRequest.MadeAt,
	}, nil
}

func (vc *VisitCreator) createBrowser() *models.Browser {
	plugins := models.Plugins{
		Java: vc.trackRequest.BrowserPluginJava,
	}

	window := models.Window{
		Width:  vc.trackRequest.BrowserWindowWidth,
		Height: vc.trackRequest.BrowserWindowHeight,
	}

	browser := models.Browser{
		Name:         vc.trackRequest.BrowserName,
		Version:      vc.trackRequest.BrowserVersion,
		MajorVersion: vc.trackRequest.BrowserMajorVersion,
		UserAgent:    vc.trackRequest.BrowserUserAgent,
		Platform:     vc.trackRequest.BrowserPlatform,
		Cookie:       vc.trackRequest.BrowserCookie,
		IsOnline:     vc.trackRequest.BrowserIsOnline,
		Plugins:      plugins,
		Window:       window,
	}

	return &browser
}

func (vc *VisitCreator) createOperatingSystem() *models.OperatingSystem {
	return &models.OperatingSystem{
		Name:    vc.trackRequest.OperatingSystemName,
		Version: vc.trackRequest.OperatingSystemVersion,
	}
}

func (vc *VisitCreator) createScreen() *models.Screen {
	return &models.Screen{
		Width:  vc.trackRequest.ScreenWidth,
		Height: vc.trackRequest.ScreenHeight,
	}
}

func (vc *VisitCreator) createDevice() *models.Device {
	return &models.Device{
		Name:     vc.trackRequest.DeviceName,
		IsTablet: vc.trackRequest.DeviceIsTablet,
		IsPhone:  vc.trackRequest.DeviceIsPhone,
		IsMobile: vc.trackRequest.DeviceIsMobile,
	}
}

func (vc *VisitCreator) createLocation() (*models.Location, error) {
	// geoLocation, err := geoIP.City(net.ParseIP(vc.trackRequest.GetRequestIP()))
	geoLocation, err := vc.geoIP.City(net.ParseIP("78.52.240.125"))

	if err != nil {
		return nil, err
	}

	return models.NewLocationFromGeoIP(geoLocation), nil
}

func (vc *VisitCreator) createPage() *models.Page {
	return &models.Page{
		Title: vc.trackRequest.PageTitle,
		Host:  vc.trackRequest.PageHost,
		URL:   vc.trackRequest.PageURL,
	}
}
