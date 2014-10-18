package repositories

import "github.com/piotrkowalczuk/gonalytics-backend/lib/models"

// VisitCollection ...
const VisitCollection = "visit"

// VisitRepository ...
type VisitRepository struct {
	Repository
}

// Insert ...
func (vr *VisitRepository) Insert(visit *models.Visit) error {
	cql := `
	INSERT INTO visits
	(
		id, ip, nb_of_actions, site_id, referrer, language, first_action_at,
		last_action_at, browser, screen, os, device, location, first_page,
		last_page
	)
	VALUES (
		?, ?, ?, ?, ?, ?, ?, ?,
		{
			name: ?, version: ?, major_version: ?, user_agent: ?, platform: ?,
			cookie: ?, is_online: ?,
			window: { width: ?, height: ? },
			plugins: { java: ? }
		},
		{ width: ?, height: ? },
		{ name: ?, version: ? },
		{ name: ?, is_mobile: ?, is_tablet: ?, is_phone: ? },
		{
			city_name: ?, city_id: ?, country_name: ?, country_code: ?,
			country_id: ?, continent_name: ?, continent_code: ?, continent_id: ?,
			latitude: ?, longitude: ?, metro_code: ?, time_zone: ?, postal_code: ?,
			is_anonymous_proxy: ?, is_satellite_provider: ?
		},
		{ title: ?, host: ?, url: ?},
		{ title: ?, host: ?, url: ?}
	)`

	return vr.Repository.Cassandra.Query(
		cql,
		visit.ID,
		visit.IP,
		visit.NbOfActions,
		visit.SiteID,
		visit.Referrer,
		visit.Language,
		visit.FirstActionAt,
		visit.LastActionAt,
		visit.Browser.Name,
		visit.Browser.Version,
		visit.Browser.MajorVersion,
		visit.Browser.UserAgent,
		visit.Browser.Platform,
		visit.Browser.Cookie,
		visit.Browser.IsOnline,
		visit.Browser.Window.Width,
		visit.Browser.Window.Height,
		visit.Browser.Plugins.Java,
		visit.Screen.Width,
		visit.Screen.Height,
		visit.OperatingSystem.Name,
		visit.OperatingSystem.Version,
		visit.Device.Name,
		visit.Device.IsMobile,
		visit.Device.IsTablet,
		visit.Device.IsPhone,
		visit.Location.CityName,
		visit.Location.CityID,
		visit.Location.CountryName,
		visit.Location.CountryCode,
		visit.Location.CountryID,
		visit.Location.ContinentName,
		visit.Location.ContinentCode,
		visit.Location.ContinentID,
		visit.Location.Latitude,
		visit.Location.Longitude,
		visit.Location.MetroCode,
		visit.Location.TimeZone,
		visit.Location.PostalCode,
		visit.Location.IsAnonymousProxy,
		visit.Location.IsSatelliteProvider,
		visit.FirstPage.Title,
		visit.FirstPage.Host,
		visit.FirstPage.URL,
		visit.LastPage.Title,
		visit.LastPage.Host,
		visit.LastPage.URL,
	).Exec()
}

// AddAction adds new action to existing map.
func (vr *VisitRepository) AddAction(action *models.Action) error {
	cql := `
		UPDATE visits SET
			actions[?] = {
				id: ?,
				visit_id: ?,
				referrer: ?,
				page: { title: ?, host: ?, url: ? },
				created_at: ?
			},
			last_action_at = ?,
			last_page = { title: ?, host: ?, url: ? }
		WHERE id = ?
	`
	return vr.Repository.Cassandra.Query(
		cql,
		action.CreatedAt,
		action.ID,
		action.VisitID,
		action.Referrer,
		action.Page.Title,
		action.Page.Host,
		action.Page.URL,
		action.CreatedAt,
		action.CreatedAt,
		action.Page.Title,
		action.Page.Host,
		action.Page.URL,
		action.VisitID,
	).Exec()
}
