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
	plan_id INT REFERENCES core.plans(id),
	created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);
CREATE TABLE IF NOT EXISTS secure.api_keys (
	id SERIAL PRIMARY KEY,
	developer_id INT REFERENCES secure.users(id),
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
