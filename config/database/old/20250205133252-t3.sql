
-- +migrate Up
CREATE TABLE "KKKK" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at timestamp DEFAULT now() NULL,
    updated_at TIMESTAMP DEFAULT now() ON UPDATE now()
);

-- +migrate Down
DROP TABLE "KKKK";
