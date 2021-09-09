CREATE TABLE IF NOT EXISTS users
(
    id                  UUID                                                           NOT NULL PRIMARY KEY,
    name                VARCHAR                                                        NOT NULL,
    login               VARCHAR UNIQUE,
    email               VARCHAR UNIQUE                                                 NOT NULL,
    pw_hash             VARCHAR,
    sub                 VARCHAR,
    roles               VARCHAR[]                                                      NOT NULL DEFAULT (ARRAY ['user']),
    image               VARCHAR,

    birthday            TIMESTAMP WITHOUT TIME ZONE,
    gender_self_defined BOOLEAN,
    gender              VARCHAR,
    location            VARCHAR,

    created_at          TIMESTAMP WITHOUT TIME ZONE DEFAULT (now() at time zone 'utc') NOT NULL,
    updated_at          TIMESTAMP WITHOUT TIME ZONE DEFAULT (now() at time zone 'utc') NOT NULL,
    deleted_at          TIMESTAMP WITHOUT TIME ZONE,

    CHECK (
            (login IS NOT NULL AND pw_hash IS NOT NULL) OR
            (sub IS NOT NULL)
        ),

    CHECK ( (gender_self_defined IS NULL AND gender IS NULL) OR
            (gender_self_defined = false AND (gender = 'male' OR gender = 'female' OR gender = 'declined')) OR
            (gender_self_defined = true AND gender IS NOT NULL AND (LENGTH(gender) > 2 AND LENGTH(gender) < 30))
        )
);