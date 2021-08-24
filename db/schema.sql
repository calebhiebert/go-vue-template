CREATE TABLE test_table
(
    id   UUID    NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL,
    data jsonb
);

CREATE TABLE test_relation
(
    id            UUID    NOT NULL PRIMARY KEY,
    name          VARCHAR NOT NULL,
    test_table_id UUID REFERENCES test_table (id)
);

CREATE TABLE users
(
    id   UUID    NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL,
    data jsonb
);