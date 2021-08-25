CREATE TABLE IF NOT EXISTS token_issuances
(
    id         UUID    NOT NULL PRIMARY KEY,
    user_id    UUID    NOT NULL REFERENCES users (id),
    ip_address VARCHAR NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (now() at time zone 'utc')
);