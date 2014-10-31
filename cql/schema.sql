DROP KEYSPACE IF EXISTS gonalytics;

CREATE KEYSPACE gonalytics WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

CREATE TABLE IF NOT EXISTS gonalytics.actions
(
    id timeuuid,
    visit_id timeuuid,
    ip varchar,
    site_id bigint,
    referrer varchar,
    language varchar,
    made_at timestamp,
    // BROWSER
    browser_name varchar,
    browser_version varchar,
    browser_major_version varchar,
    browser_user_agent varchar,
    browser_platform varchar,
    browser_cookie boolean,
    browser_is_online boolean,
    browser_window_width int,
    browser_window_height int,
    browser_plugin_java boolean,
    // SCREEN
    screen_width int,
    screen_height int,
    // OS
    os_name varchar,
    os_version varchar,
    // DEVICE
    device_name varchar,
    device_is_mobile boolean,
    device_is_tablet boolean,
    device_is_phone boolean,
    // LOCATION
    location_city_name varchar,
    location_city_id int,
    location_country_name varchar,
    location_country_code varchar,
    location_country_id int,
    location_continent_name varchar,
    location_continent_code varchar,
    location_continent_id int,
    location_latitude double,
    location_longitude double,
    location_time_zone varchar,
    location_metro_code int,
    location_postal_code varchar,
    location_is_anonymous_proxy boolean,
    location_is_satellite_provider boolean,
    // PAGE
    page_title varchar,
    page_host varchar,
    page_url varchar,
    PRIMARY KEY (id),
) WITH comment='Column family contains actions.';
