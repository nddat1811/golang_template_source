
-- +migrate Up

CREATE TABLE example_table2 (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    t1_id INT,
    CONSTRAINT fk_customer
   FOREIGN KEY(t1_id)
      REFERENCES example_table(id)
);
-- +migrate Down
DROP TABLE IF EXISTS example_table2;