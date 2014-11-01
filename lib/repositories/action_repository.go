package repositories

import (
	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
	"github.com/relops/cqlr"
)

const (
	// ActionColumnFamily ...
	ActionColumnFamily = "actions"
	allFields          = `
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
	INSERT INTO ` + ActionColumnFamily + `
	(
		` + allFields + `
	)
	VALUES (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	)`

	return cqlr.Bind(cql, action).Exec(ar.Repository.Cassandra)
}

// Find ...
func (ar *ActionRepository) Find() ([]*models.ActionEntity, error) {
	cql := `SELECT ` + allFields + ` FROM ` + ActionColumnFamily

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
