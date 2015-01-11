DROP KEYSPACE IF EXISTS gonalytics;

CREATE KEYSPACE gonalytics WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

CREATE TABLE IF NOT EXISTS gonalytics.visit_actions
(
    id timeuuid,
    visit_id timeuuid,
    ip varchar,
    site_id bigint,
    referrer varchar,
    language varchar,
    // MADE AT
    made_at timestamp,
    made_at_year int,
    made_at_month int,
    made_at_week int,
    made_at_day int,
    made_at_hour int,
    made_at_minute int,
    made_at_second int,
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
    PRIMARY KEY (visit_id, made_at),
) WITH comment='Column family contains actions.';

CREATE TABLE IF NOT EXISTS gonalytics.site_day_country_actions_counter
(
    site_id bigint,
    count counter,
    // LOCATION
    location_country_name varchar,
    location_country_code varchar,
    location_country_id int,
    // MADE AT
    made_at_year int,
    made_at_month int,
    made_at_day int,

    PRIMARY KEY ((site_id, made_at_year, made_at_month, made_at_day), location_country_name, location_country_code, location_country_id),
);

CREATE TABLE IF NOT EXISTS gonalytics.site_month_country_actions_counter
(
    site_id bigint,
    count counter,
    // LOCATION
    location_country_name varchar,
    location_country_code varchar,
    location_country_id int,
    // MADE AT
    made_at_year int,
    made_at_month int,
    
    PRIMARY KEY ((site_id, made_at_year, made_at_month), location_country_name, location_country_code, location_country_id),
);

CREATE TABLE IF NOT EXISTS gonalytics.site_year_country_actions_counter
(
    site_id bigint,
    count counter,
    // LOCATION
    location_country_name varchar,
    location_country_code varchar,
    location_country_id int,
    // MADE AT
    made_at_year int,
    
    PRIMARY KEY ((site_id, made_at_year), location_country_name, location_country_code, location_country_id),
);

CREATE TABLE IF NOT EXISTS gonalytics.site_day_browser_actions_counter
(
    site_id bigint,
    count counter,
    // BROWSER
    browser_name varchar,
    browser_version varchar,
    // MADE AT
    made_at_year int,
    made_at_month int,
    made_at_day int,

    PRIMARY KEY ((site_id, made_at_year, made_at_month, made_at_day), browser_name, browser_version),
);

CREATE TABLE IF NOT EXISTS gonalytics.site_month_browser_actions_counter
(
    site_id bigint,
    count counter,
    // BROWSER
    browser_name varchar,
    browser_version varchar,
    // MADE AT
    made_at_year int,
    made_at_month int,
    
    PRIMARY KEY ((site_id, made_at_year, made_at_month), browser_name, browser_version),
);

CREATE TABLE IF NOT EXISTS gonalytics.site_year_browser_actions_counter
(
    site_id bigint,
    count counter,
    // BROWSER
    browser_name varchar,
    browser_version varchar,
    // MADE AT
    made_at_year int,
    
    PRIMARY KEY ((site_id, made_at_year), browser_name, browser_version),
);


CREATE TABLE IF NOT EXISTS gonalytics.site_day_country_visits_counter
(
    site_id bigint,
    count counter,
    // LOCATION
    location_country_name varchar,
    location_country_code varchar,
    location_country_id int,
    // MADE AT
    made_at_year int,
    made_at_month int,
    made_at_day int,
    
    PRIMARY KEY ((site_id, made_at_year, made_at_month, made_at_day), location_country_name, location_country_code, location_country_id),
);

CREATE TABLE IF NOT EXISTS gonalytics.site_month_country_visits_counter
(
    site_id bigint,
    count counter,
    // LOCATION
    location_country_name varchar,
    location_country_code varchar,
    location_country_id int,
    // MADE AT
    made_at_year int,
    made_at_month int,
    
    PRIMARY KEY ((site_id, made_at_year, made_at_month), location_country_name, location_country_code, location_country_id),
);

CREATE TABLE IF NOT EXISTS gonalytics.site_year_country_visits_counter
(
    site_id bigint,
    count counter,
    // LOCATION
    location_country_name varchar,
    location_country_code varchar,
    location_country_id int,
    // MADE AT
    made_at_year int,
    
    PRIMARY KEY ((site_id, made_at_year), location_country_name, location_country_code, location_country_id),
);

CREATE TABLE IF NOT EXISTS gonalytics.site_day_browser_visits_counter
(
    site_id bigint,
    count counter,
    // BROWSER
    browser_name varchar,
    browser_version varchar,
    // MADE AT
    made_at_year int,
    made_at_month int,
    made_at_day int,

    PRIMARY KEY ((site_id, made_at_year, made_at_month, made_at_day), browser_name, browser_version),
);

CREATE TABLE IF NOT EXISTS gonalytics.site_month_browser_visits_counter
(
    site_id bigint,
    count counter,
    // BROWSER
    browser_name varchar,
    browser_version varchar,
    // MADE AT
    made_at_year int,
    made_at_month int,
    
    PRIMARY KEY ((site_id, made_at_year, made_at_month), browser_name, browser_version),
);

CREATE TABLE IF NOT EXISTS gonalytics.site_year_browser_visits_counter
(
    site_id bigint,
    count counter,
    // BROWSER
    browser_name varchar,
    browser_version varchar,
    // MADE AT
    made_at_year int,
    
    PRIMARY KEY ((site_id, made_at_year), browser_name, browser_version),
);