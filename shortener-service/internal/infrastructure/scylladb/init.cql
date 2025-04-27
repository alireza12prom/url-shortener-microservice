CREATE KEYSPACE IF NOT EXISTS shortener
WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

USE shortener;

CREATE TABLE IF NOT EXISTS account (
  id UUID,
  name VARCHAR,
  username VARCHAR,
  password VARCHAR,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,

  PRIMARY KEY(id)
);

CREATE MATERIALIZED VIEW account_by_username AS
  SELECT * FROM account WHERE username IS NOT NULL
  PRIMARY KEY (username, id);