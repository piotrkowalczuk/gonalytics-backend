DROP KEYSPACE IF EXISTS gonalytics;

CREATE KEYSPACE gonalytics WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

CREATE TYPE gonalytics.window (
   width int,
   height int
);

CREATE TYPE gonalytics.plugins (
   java boolean
);

CREATE TYPE gonalytics.browser (
   name varchar,
   version varchar,
   major_version varchar,
   user_agent varchar,
   platform varchar,
   cookie boolean,
   is_online boolean,
   window frozen<gonalytics.window>,
   plugins frozen<gonalytics.plugins>
);

CREATE TYPE gonalytics.screen (
    width int,
    height int
);

CREATE TYPE gonalytics.operating_system (
    name varchar,
    version varchar
);

CREATE TYPE gonalytics.device (
    name varchar,
    is_mobile boolean,
    is_tablet boolean,
    is_phone boolean
);

CREATE TYPE gonalytics.location (
    city_name varchar,
    city_id int,
    country_name varchar,
    country_code varchar,
    country_id int,
    continent_name varchar,
    continent_code varchar,
    continent_id int,
    latitude double,
    longitude double,
    time_zone varchar,
    metro_code int,
    postal_code varchar,
    is_anonymous_proxy boolean,
    is_satellite_provider boolean
);

CREATE TYPE gonalytics.page (
    title varchar,
    host varchar,
    url varchar
);

CREATE TYPE gonalytics.action (
    id timeuuid,
    visit_id timeuuid,
    referrer varchar,
    page frozen<gonalytics.page>,
    created_at timestamp
);

CREATE TABLE IF NOT EXISTS gonalytics.visits
(
    id timeuuid,
    ip varchar,
    nb_of_actions bigint,
    site_id bigint,
    referrer varchar,
    language varchar,
    browser frozen<gonalytics.browser>,
    screen frozen<gonalytics.screen>,
    os frozen<gonalytics.operating_system>,
    device frozen<gonalytics.device>,
    location frozen<gonalytics.location>,
    first_page frozen<gonalytics.page>,
    last_page frozen<gonalytics.page>,
    actions map<timestamp, frozen<gonalytics.action>>,
    first_action_at timestamp,
    last_action_at timestamp,
    PRIMARY KEY (id)
) WITH comment='Visits and related actions.';
