CREATE TABLE IF NOT EXISTS users
(
    id         UUID                                                           NOT NULL PRIMARY KEY,
    name       VARCHAR                                                        NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (now() at time zone 'utc') NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (now() at time zone 'utc') NOT NULL
);