
-- +migrate Up

CREATE TABLE example_table (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);
-- +migrate Down
DROP TABLE IF EXISTS example_table;