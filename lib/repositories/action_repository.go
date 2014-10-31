package repositories

import (
	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
	"github.com/relops/cqlr"
)

const (
	// ActionCollection ...
	ActionCollection = "actions"
	allFields        = `
		id, ip, visit_id, site_id, referrer, language, browser_name,
		browser_version, browser_major_version, browser_user_agent,
		browser_platform, browser_cookie, browser_plugin_java, browser_is_online,
		browser_window_width, browser_window_height, screen_width, screen_height,
		os_name, os_version, device_name, device_is_mobile, device_is_tablet,
		device_is_phone, location_city_name, location_city_id,
		location_country_name, location_country_code, location_country_id,
		location_continent_name, location_continent_code, location_continent_id,
		location_latitude, location_longitude, location_metro_code,
		location_time_zone, location_postal_code, location_is_anonymous_proxy,
		location_is_satellite_provider, page_title, page_host, page_url,
		made_at
	`
)

// ActionRepository ...
type ActionRepository struct {
	Repository
}

// Insert ...
func (ar *ActionRepository) Insert(action *models.ActionEntity) error {
	cql := `
	INSERT INTO actions
	(
		` + allFields + `
	)
	VALUES (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	)`

	return cqlr.Bind(cql, action).Exec(ar.Repository.Cassandra)
	// return vr.Repository.Cassandra.Query(
	// 	cql,
	// 	visit.ID,
	// 	visit.IP,
	// 	visit.NbOfActions,
	// 	visit.SiteID,
	// 	visit.Referrer,
	// 	visit.Language,
	// 	visit.FirstActionAt,
	// 	visit.LastActionAt,
	// 	visit.Browser.Name,
	// 	visit.Browser.Version,
	// 	visit.Browser.MajorVersion,
	// 	visit.Browser.UserAgent,
	// 	visit.Browser.Platform,
	// 	visit.Browser.Cookie,
	// 	visit.Browser.IsOnline,
	// 	visit.Browser.Window.Width,
	// 	visit.Browser.Window.Height,
	// 	visit.Browser.Plugins.Java,
	// 	visit.Screen.Width,
	// 	visit.Screen.Height,
	// 	visit.OperatingSystem.Name,
	// 	visit.OperatingSystem.Version,
	// 	visit.Device.Name,
	// 	visit.Device.IsMobile,
	// 	visit.Device.IsTablet,
	// 	visit.Device.IsPhone,
	// 	visit.Location.CityName,
	// 	visit.Location.CityID,
	// 	visit.Location.CountryName,
	// 	visit.Location.CountryCode,
	// 	visit.Location.CountryID,
	// 	visit.Location.ContinentName,
	// 	visit.Location.ContinentCode,
	// 	visit.Location.ContinentID,
	// 	visit.Location.Latitude,
	// 	visit.Location.Longitude,
	// 	visit.Location.MetroCode,
	// 	visit.Location.TimeZone,
	// 	visit.Location.PostalCode,
	// 	visit.Location.IsAnonymousProxy,
	// 	visit.Location.IsSatelliteProvider,
	// 	visit.FirstPage.Title,
	// 	visit.FirstPage.Host,
	// 	visit.FirstPage.URL,
	// 	visit.LastPage.Title,
	// 	visit.LastPage.Host,
	// 	visit.LastPage.URL,
	// ).Exec()
}

// Find ...
func (ar *ActionRepository) Find() ([]*models.ActionEntity, error) {
	cql := `SELECT ` + allFields + ` FROM  actions`

	query := ar.Repository.Cassandra.Query(cql, "me").Consistency(gocql.One)

	b := cqlr.BindQuery(query)

	actions := []*models.ActionEntity{}
	var a models.ActionEntity

	for b.Scan(&a) {
		actions = append(actions, &a)
	}

	if err := b.Close(); err != nil {
		return nil, err
	}

	return actions, nil
}
