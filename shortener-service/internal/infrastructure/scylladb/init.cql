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

CREATED TABLE IF NOT EXISTS short_url {
  id UUID,
  user_id UUID,
  endpoint VARCHAR,
  hash VARCHAR,
  is_active BOOL,
  created_at TIMESTAMP,
  last_accessed_at TIMESTAMP,

  PRIMARY KEY (id)
}

CREATE MATERIALIZED VIEW short_url_by_user_id AS
  SELECT * FROM short_url WHERE user_id IS NOT NULL
  PRIMARY KEY (user_id, id);

CREATE MATERIALIZED VIEW short_url_by_hash AS
  SELECT * FROM short_url WHERE hash IS NOT NULL
  PRIMARY KEY (hash, id);