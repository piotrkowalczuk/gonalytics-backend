CREATE KEYSPACE gonalytics WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

CREATE TABLE visits
(
    id timeuuid,
    ip varchar,
    nb_of_actions bigint,
    site_id bigint,
    referrer varchar,
    language varchar,
    first_action_at timestamp,
    last_action_at timestamp,
    PRIMARY KEY (id)
);
