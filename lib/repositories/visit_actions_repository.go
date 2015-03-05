package repositories

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
	"github.com/relops/cqlr"
)

const (
	// VisitActionsColumnFamily ...
	VisitActionsColumnFamily = "visit_actions"
	// VisitActionsFields ...
	VisitActionsFields = `
		id, ip, visit_id, site_id, referrer, language, browser_name,
		browser_version, browser_major_version, browser_user_agent,
		browser_platform, browser_cookie, browser_plugin_java, browser_is_online,
		browser_window_width, browser_window_height, screen_width, screen_height,
		os_name, os_version, device_name, device_is_mobile, device_is_tablet,
		device_is_phone, location_city_name, location_city_id,
		location_country_name, location_country_code, location_country_id,
		location_continent_name, location_continent_code, location_continent_id,
		location_latitude, location_longitude, location_metro_code,
		location_time, location_timezone, location_postal_code,
		location_is_anonymous_proxy,
		location_is_satellite_provider, page_title, page_host, page_url,
		server_time
	`
)

// VisitActionsRepository ...
type VisitActionsRepository struct {
	Repository
	ColumnFamily string
}

// Insert ...
func (r *VisitActionsRepository) Insert(action *models.ActionEntity) error {
	cql := `
	INSERT INTO ` + VisitActionsColumnFamily + `
	(
		` + VisitActionsFields + `
	)
	VALUES (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	)`

	return cqlr.Bind(cql, action).Exec(r.Repository.Cassandra)
}

// IsActiveVisit ...
func (r *VisitActionsRepository) IsActiveVisit(visitID gocql.UUID) (bool, error) {
	nbOfActions := 0

	cql := `SELECT COUNT(*) FROM ` + VisitActionsColumnFamily +
		` WHERE visit_id = ? AND server_time >= ? LIMIT 1`

	iter := r.Repository.Cassandra.Query(
		cql,
		visitID,
		time.Now().UTC().Add(-models.MinVisitDuration),
	).Consistency(gocql.One).
		Iter()

	iter.Scan(&nbOfActions)

	if err := iter.Close(); err != nil {
		return false, err
	}

	return nbOfActions == 1, nil
}

// Find ...
func (r *VisitActionsRepository) Find() ([]*models.ActionEntity, error) {
	cql := `SELECT ` + VisitActionsFields + ` FROM ` + VisitActionsColumnFamily

	query := r.Repository.Cassandra.Query(cql).Consistency(gocql.One)

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
