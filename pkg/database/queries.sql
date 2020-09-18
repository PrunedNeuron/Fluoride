-- First, create the required schemas
CREATE SCHEMA IF NOT EXISTS core AUTHORIZATION ayush;
CREATE SCHEMA IF NOT EXISTS secure AUTHORIZATION ayush;
CREATE SCHEMA IF NOT EXISTS icon_packs AUTHORIZATION ayush;
CREATE SCHEMA IF NOT EXISTS icon_requests AUTHORIZATION ayush;


CREATE TABLE IF NOT EXISTS core.plans (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	subtitle TEXT,
	description TEXT NOT NULL,
	price NUMERIC NOT NULL,
	billing_cycle NUMERIC NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);
CREATE TABLE IF NOT EXISTS secure.users (
	id SERIAL PRIMARY KEY,
	role TEXT NOT NULL,
	name TEXT NOT NULL,
	username TEXT UNIQUE NOT NULL,
	email TEXT UNIQUE NOT NULL,
	url TEXT,
	created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);
CREATE TABLE IF NOT EXISTS secure.billings (
	id SERIAL PRIMARY KEY,
	developer_id INT REFERENCES secure.users(id),
	plan_id INT REFERENCES plans(id),
	created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);
CREATE TABLE IF NOT EXISTS secure.api_keys (
	id SERIAL PRIMARY KEY,
	developer_id INT REFERENCES secure.users(id),
	icon_pack_id INT REFERENCES icon_packs.icon_packs(id),
	api_key TEXT UNIQUE NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);
CREATE TABLE IF NOT EXISTS secure.credentials (
	id SERIAL PRIMARY KEY,
	username TEXT REFERENCES secure.users(username),
	password TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);
/* CREATE TABLE IF NOT EXISTS icon_packs.icon_packs (
	id SERIAL PRIMARY KEY,
	name TEXT UNIQUE NOT NULL,
	developer_username TEXT REFERENCES secure.users(username),
	url TEXT NOT NULL,
	billing_status TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);
CREATE TABLE IF NOT EXISTS icon_requests.icon_requests (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	component TEXT UNIQUE NOT NULL,
	url TEXT NOT NULL,
	requesters TEXT NOT NULL,
	status TEXT NOT NULL,
	icon_pack_name INT REFERENCES icon_packs.icon_packs(name)
	created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
); */
