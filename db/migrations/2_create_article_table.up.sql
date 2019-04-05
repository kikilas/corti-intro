CREATE TABLE IF NOT EXISTS article (
	id              SERIAL PRIMARY KEY,
	text            STRING NOT NULL,
	author_id       INT
);