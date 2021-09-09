\set ON_ERROR_STOP 1

DROP DATABASE IF EXISTS mdb;
DROP user IF EXISTS mdb;
CREATE DATABASE mdb;
CREATE user mdb WITH PASSWORD 'mdb';

\connect mdb

CREATE SCHEMA mdb;
GRANT usage ON SCHEMA mdb TO mdb;

create table mdb.users
(
    id varchar(100) PRIMARY KEY,
    funds integer not null
    constraint funds_nonnegative check (funds >= 0)
);

GRANT SELECT, INSERT, UPDATE, DELETE ON mdb.users TO mdb;
COMMENT ON TABLE mdb.users IS 'Пользователи';



create table mdb.transactions
(
    id serial PRIMARY KEY,
    user_id_from varchar(100),
    user_id_to   varchar(100),
    comment varchar(500),
    creation_date timestamp NOT NULL DEFAULT NOW(),
    funds  integer not null
);

GRANT SELECT, INSERT, UPDATE, DELETE ON mdb.transactions TO mdb;
COMMENT ON TABLE mdb.transactions IS 'Транзакции';

GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA mdb TO mdb;
-- create table mdb.funds
-- (
--     id varchar(100) REFERENCES mdb.users(id) ON DELETE CASCADE,

-- );
--
-- GRANT SELECT, INSERT, UPDATE, DELETE ON mdb.funds TO mdb;
-- COMMENT ON TABLE mdb.funds IS 'Баланс';
--
-- INSERT INTO mdb.users (id, name)
-- VALUES ('admin', 'some name'),
-- VALUES ('test1', 'noname1'),
-- VALUES ('test2', 'noname2'),
-- VALUES ('test3', 'noname3'),
-- VALUES ('test4', 'noname4');