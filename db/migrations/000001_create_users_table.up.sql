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
    phone               VARCHAR UNIQUE,

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

insert into public.users (id, name, login, email, pw_hash, sub, roles, image, birthday, gender_self_defined, gender,
                          location)
values ('7403c893-cd37-41b5-85a5-a3adc51f138e', 'Admin', 'admin@admin.admin', 'admin@admin.admin',
        '$2a$14$uYJgftrFx/tXBst6w8IdaeM7MTbfajUibXmfIm4X.VobeKia.ub5m', null, '{user,admin}', null, null, null, null,
        null),
       ('54edce03-c535-4342-8518-4d5b9b1194b6', 'User', 'user@user.user', 'user@user.user',
        '$2a$14$HOrBkLTNcyDA98AC./YhEu/p5YBlKV/UAD6.1jJf2oOGpz6UVSwau', null, '{user}', null, null, null, null, null);
