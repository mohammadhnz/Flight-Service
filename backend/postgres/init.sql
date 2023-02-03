CREATE DATABASE project1;
\c project1;
CREATE USER user1 WITH PASSWORD 'password1';
GRANT ALL PRIVILEGES ON DATABASE "project1" to user1;
GRANT ALL ON SCHEMA public TO user1;
\c - user1


CREATE TABLE IF NOT EXISTS user_account
(
    user_id       SERIAL PRIMARY KEY,
    email         VARCHAR UNIQUE NOT NULL,
    phone_number  VARCHAR UNIQUE NOT NULL,
    gender        VARCHAR(1),
    first_name    VARCHAR,
    last_name     VARCHAR,
    password_hash VARCHAR
);

CREATE UNIQUE INDEX ON user_account (email);

CREATE UNIQUE INDEX ON user_account (phone_number);

CREATE TABLE IF NOT EXISTS unauthorized_token
(
    user_id    INTEGER REFERENCES user_account ON DELETE CASCADE ON UPDATE CASCADE,
    token      VARCHAR,
    expiration TIMESTAMP
);

CREATE INDEX ON unauthorized_token (token);




-- Aircrafts

CREATE TABLE IF NOT EXISTS aircraft_type
(
    type_id      VARCHAR PRIMARY KEY,
    manufacturer VARCHAR NOT NULL,
    model        VARCHAR NOT NULL,
    series       VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS aircraft_layout
(
    layout_id        SERIAL PRIMARY KEY,
    type_id          VARCHAR NOT NULL REFERENCES aircraft_type ON DELETE CASCADE ON UPDATE CASCADE,
    y_class_capacity INTEGER NOT NULL,
    j_class_capacity INTEGER NOT NULL,
    f_class_capacity INTEGER NOT NULL,
    UNIQUE (type_id, y_class_capacity, j_class_capacity, f_class_capacity)
);

CREATE TABLE IF NOT EXISTS aircraft
(
    registration VARCHAR(6) NOT NULL PRIMARY KEY,
    layout_id    INTEGER    NOT NULL REFERENCES aircraft_layout ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE VIEW aircraft_view AS
SELECT aircraft.registration                                                                   as registration,
       CONCAT(aircraft_type.manufacturer, ' ', aircraft_type.model, '-', aircraft_type.series) as aircraft_type,
       aircraft_type.type_id                                                                   as type_id,
       aircraft_layout.y_class_capacity                                                        as y_class_capacity,
       aircraft_layout.j_class_capacity                                                        as j_class_capacity,
       aircraft_layout.f_class_capacity                                                        as f_class_capacity
FROM aircraft_type
         JOIN aircraft_layout on aircraft_type.type_id = aircraft_layout.type_id
         JOIN aircraft on aircraft_layout.layout_id = aircraft.layout_id;

-- Locations

CREATE TABLE IF NOT EXISTS country
(
    country_name VARCHAR PRIMARY KEY
);

-- https://www.postgresql.org/docs/8.1/datetime-keywords.html
CREATE TABLE IF NOT EXISTS city
(
    country_name  VARCHAR REFERENCES country ON DELETE RESTRICT ON UPDATE RESTRICT,
    city_name     VARCHAR,
    timezone_name VARCHAR NOT NULL,
    PRIMARY KEY (country_name, city_name)
);

CREATE TABLE IF NOT EXISTS airport
(
    country_name VARCHAR NOT NULL,
    city_name    VARCHAR NOT NULL,
    airport_name VARCHAR NOT NULL,
    iata_code    VARCHAR(3) PRIMARY KEY,
    FOREIGN KEY (country_name, city_name) REFERENCES city (country_name, city_name) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE INDEX ON airport (country_name, city_name);

CREATE VIEW origin_destination AS
SELECT country_name as county,
       city_name    as city,
       airport_name as airport,
       iata_code    as iata
FROM airport
UNION
SELECT country_name   as county,
       city_name      as city,
       'All airports' as airport,
       'ALL'          as iata
FROM city;

CREATE VIEW airport_timezone AS
SELECT airport.iata_code, city.timezone_name
FROM airport
         JOIN city ON city.country_name = airport.country_name AND city.city_name = airport.city_name;

-- Flights

CREATE TABLE IF NOT EXISTS flight
(
    flight_serial SERIAL PRIMARY KEY,
    flight_id     VARCHAR    NOT NULL,
    origin        VARCHAR(3) NOT NULL REFERENCES airport ON DELETE CASCADE ON UPDATE CASCADE,
    destination   VARCHAR(3) NOT NULL REFERENCES airport ON DELETE CASCADE ON UPDATE CASCADE,
    aircraft      VARCHAR(6) NOT NULL REFERENCES aircraft ON DELETE CASCADE ON UPDATE CASCADE,
    departure_utc TIMESTAMP  NOT NULL,
    duration      INTERVAL   NOT NULL,
    y_price       INTEGER    NOT NULL,
    j_price       INTEGER    NOT NULL,
    f_price       INTEGER    NOT NULL
);

CREATE INDEX ON flight (flight_id);

CREATE INDEX ON flight (origin, destination, departure_utc);

-- Transactions

-- Purchases

CREATE TABLE IF NOT EXISTS purchase
(
    corresponding_user_id INTEGER,
    title                 VARCHAR,
    first_name            VARCHAR,
    last_name             VARCHAR,
    flight_serial         INTEGER,
    offer_price           INTEGER,
    offer_class           VARCHAR
    -- TODO: transactions
);

-- OFFERS

CREATE VIEW available_offers AS
SELECT flight.flight_id                                        AS flight_id,
       flight.origin                                           AS origin,
       flight.destination                                      AS destination,
       flight.departure_utc::TIMESTAMP WITH TIME ZONE AT TIME ZONE
       (SELECT timezone_name
        FROM airport_timezone
        WHERE airport_timezone.iata_code = flight.origin)      AS departure_local_time,
       (flight.departure_utc + flight.duration)::TIMESTAMP WITH TIME ZONE AT TIME ZONE
       (SELECT timezone_name
        FROM airport_timezone
        WHERE airport_timezone.iata_code = flight.destination) AS arrival_local_time,
       flight.duration                                         AS duration,
       flight.y_price                                          AS y_price,
       flight.j_price                                          AS j_price,
       flight.f_price                                          AS f_price,
       ((SELECT y_class_capacity FROM aircraft_view WHERE aircraft_view.registration = flight.aircraft) -
        (SELECT COUNT(*)
         FROM purchase
         WHERE purchase.flight_serial = flight.flight_serial
           AND offer_class = 'Y'))                             AS y_class_free_capacity,
       ((SELECT j_class_capacity FROM aircraft_view WHERE aircraft_view.registration = flight.aircraft) -
        (SELECT COUNT(*)
         FROM purchase
         WHERE purchase.flight_serial = flight.flight_serial
           AND offer_class = 'J'))
                                                               AS j_class_free_capacity,
       ((SELECT f_class_capacity FROM aircraft_view WHERE aircraft_view.registration = flight.aircraft) -
        (SELECT COUNT(*)
         FROM purchase
         WHERE purchase.flight_serial = flight.flight_serial
           AND offer_class = 'F'))                             AS f_class_free_capacity,
       aircraft_view.aircraft_type                             AS equipment
FROM flight
         JOIN aircraft_view on aircraft_view.registration = flight.aircraft;


\COPY aircraft_type (type_id,manufacturer,model,series) FROM '/var/lib/postgresql/csvs/aircraft_type.csv' DELIMITER ',' CSV HEADER;   
\COPY aircraft_layout (layout_id,type_id,y_class_capacity,j_class_capacity,f_class_capacity) FROM '/var/lib/postgresql/csvs/aircraft_layout.csv' DELIMITER ',' CSV HEADER;         
\COPY aircraft (registration,layout_id) FROM '/var/lib/postgresql/csvs/aircraft.csv' DELIMITER ',' CSV HEADER;
\COPY country (country_name) FROM '/var/lib/postgresql/csvs/country.csv' DELIMITER ',' CSV HEADER;
\COPY city (country_name,city_name,timezone_name) FROM '/var/lib/postgresql/csvs/city.csv' DELIMITER ',' CSV HEADER;
\COPY airport (country_name,city_name,airport_name,iata_code) FROM '/var/lib/postgresql/csvs/airport.csv' DELIMITER ',' CSV HEADER;
\COPY flight (flight_serial,flight_id,origin,destination,aircraft,departure_utc,duration,y_price,j_price,f_price) FROM '/var/lib/postgresql/csvs/flight.csv' DELIMITER ',' CSV HEADER;
