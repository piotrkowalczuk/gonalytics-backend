package services

import (
	"github.com/piotrkowalczuk/gonalytics-tracker/models"
	"github.com/piotrkowalczuk/gonalytics-tracker/structs"
	"labix.org/v2/mgo/bson"
)

// VisitCreator ...
type VisitCreator struct {
	trackRequest *structs.TrackRequest
	Visit        *models.Visit
}

// NewVisitCreator ...
func NewVisitCreator(trackRequest *structs.TrackRequest) *VisitCreator {
	vc := VisitCreator{
		trackRequest: trackRequest,
	}

	vc.createVisit()

	return &vc
}

func (vc *VisitCreator) createVisit() {
	vc.Visit = &models.Visit{
		ID:                  bson.NewObjectId(),
		IP:                  vc.trackRequest.GetRequestIP(),
		Referrer:            vc.trackRequest.Referrer,
		Language:            vc.trackRequest.Language,
		NbOfActions:         1,
		SiteID:              vc.trackRequest.SiteID,
		Location:            vc.createLocation(),
		Browser:             vc.createBrowser(),
		FirstPage:           vc.createPage(),
		LastPage:            vc.createPage(),
		OperatingSystem:     vc.createOperatingSystem(),
		Screen:              vc.createScreen(),
		Device:              vc.createDevice(),
		FirstActionAt:       vc.trackRequest.MadeAt,
		FirstActionAtBucket: vc.trackRequest.MadeAtBucket,
		LastActionAt:        vc.trackRequest.MadeAt,
		LastActionAtBucket:  vc.trackRequest.MadeAtBucket,
	}
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

func (vc *VisitCreator) createLocation() *models.Location {
	geoLocation, err := NewGeoLocation(vc.trackRequest.GetRequestIP())

	location := &models.Location{}
	if err == nil {
		location = models.NewLocationFromGeoIP(geoLocation.Location)
	}

	return location
}

func (vc *VisitCreator) createPage() *models.Page {
	return &models.Page{
		Title: vc.trackRequest.PageTitle,
		Host:  vc.trackRequest.PageHost,
		Url:   vc.trackRequest.PageURL,
	}
}
