CREATE TABLE IF NOT EXISTS access_logs
(
    id                  UUID    NOT NULL PRIMARY KEY,
    path                VARCHAR NOT NULL,
    request_body        TEXT,
    request_headers     JSONB,
    response_body       JSONB   NOT NULL,
    response_headers    JSONB   NOT NULL,
    response_code       INTEGER NOT NULL,
    processing_duration INTEGER NOT NULL,
    user_id             UUID REFERENCES users (id),
    ip_address          VARCHAR NOT NULL,
    created_at          TIMESTAMP WITHOUT TIME ZONE DEFAULT (now() at time zone 'utc')
);