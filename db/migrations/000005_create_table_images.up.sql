CREATE TABLE IF NOT EXISTS images
(
    id         UUID                                                           NOT NULL PRIMARY KEY,
    name       VARCHAR                                                        NOT NULL,

    type       VARCHAR                                                        NOT NULL,
    size       INTEGER                                                        NOT NULL,

    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (now() at time zone 'utc') NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (now() at time zone 'utc') NOT NULL,
    deleted_at TIMESTAMP WITHOUT TIME ZONE
);