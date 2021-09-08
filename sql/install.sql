\set ON_ERROR_STOP 1

DROP DATABASE IF EXISTS mdb;
DROP user IF EXISTS mdb;
CREATE DATABASE mdb;
CREATE user mdb WITH PASSWORD 'mdb';

\connect mdb

CREATE SCHEMA mdb;
GRANT usage ON SCHEMA mdb TO mdb;

create table mdb.user
(
    id varchar(100) PRIMARY KEY,
    funds integer not null
    constraint funds_nonnegative check (funds >= 0)
);

GRANT SELECT, INSERT, UPDATE, DELETE ON mdb.user TO mdb;
COMMENT ON TABLE mdb.user IS 'Пользователь';