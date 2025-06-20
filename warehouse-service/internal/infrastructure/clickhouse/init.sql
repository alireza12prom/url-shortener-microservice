CREATE TABLE domain_events (
    id UUID,
    name String,
    context String,
    payload JSON,
    datetime DateTime,
    correlation_id UUID
) ENGINE = MergeTree
ORDER BY (datetime);