CREATE TABLE policys (
	id SERIAL PRIMARY KEY,
	name VARCHAR(128),
	manifest TEXT,
	accountId INTEGER,
	created timestamp(0) without time zone default (now() at time zone 'utc')
);

