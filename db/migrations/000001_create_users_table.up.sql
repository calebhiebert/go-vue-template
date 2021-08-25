CREATE TABLE IF NOT EXISTS users
(
    id         UUID                                                           NOT NULL PRIMARY KEY,
    name       VARCHAR                                                        NOT NULL,
    login      VARCHAR UNIQUE,
    email      VARCHAR UNIQUE                                                 NOT NULL,
    pw_hash    VARCHAR,
    sub        VARCHAR,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (now() at time zone 'utc') NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (now() at time zone 'utc') NOT NULL,
    deleted_at TIMESTAMP WITHOUT TIME ZONE,

    CHECK (
            (login IS NOT NULL AND pw_hash IS NOT NULL) OR
            (sub IS NOT NULL)
        )
);