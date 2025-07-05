CREATE DATABASE IF NOT EXISTS warehouse;

CREATE TABLE IF NOT EXISTS warehouse.domain_events (
    id UUID,
    name String,
    context String,
    payload JSON,
    datetime DateTime,
    correlation_id UUID
) ENGINE = MergeTree
ORDER BY (datetime);