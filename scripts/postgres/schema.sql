create type wind_dir_type as enum('N', 'NNE', 'NE', 'ENE', 'E', 'ESE', 'SE', 'SSE', 'S', 'SSW', 'SW', 'WSW', 'W', 'WNW', 'NW', 'NNW')

CREATE TABLE weather(
    id BIGSERIAL PRIMARY KEY,
    last_updated_epoch bigint NOT NULL,
	last_updated       timestamp with time zone NOT NULL,
	temp_c             real,
	temp_f             real,
	is_day             bigint,
	condition          text NOT NULL,
	wind_mph           real,
	wind_kph           real,
	wind_degree        bigint,
	wind_dir           wind_dir_type NOT NULL,
	pressure_mb        real,
	pressure_in        real,
	precip_mm          real,
	precip_in          real,
	humidity           bigint,
	cloud              bigint,
	feelslike_c        real,
	feelslike_f        real,
	vis_km             real,
	vis_miles          real,
	uv                 real,
	gust_mph           real,
	gust_kph           real
);

CREATE TABLE location_info (
	id 				BIGSERIAL PRIMARY KEY,
	name            text NOT NULL,
	region          text NOT NULL,
	country         text NOT NULL,
	lat             real, 
	lon             real,
	tz_id           text NOT NULL,
	localtime_epoch bigint,
	localtime       timestamp with time zone NOT NULL,
	UNIQUE (name, region, country, lat, lon, tz_id)
);