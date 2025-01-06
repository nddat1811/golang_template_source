-- Create database if it doesn't exist (PostgreSQL doesn't support IF NOT EXISTS in CREATE DATABASE).
-- You can run this command only once to create the database.
-- CREATE DATABASE test;

-- Switch to the test database
-- In PostgreSQL, you switch databases using a separate connection, so skip the USE command.

-- Set PostgreSQL-specific settings (optional)
SET client_encoding = 'UTF8';
SET timezone = 'UTC';

-- Drop the `users` table if it exists
DROP TABLE IF EXISTS users;

-- Create the `users` table
CREATE TABLE users (
    id SERIAL PRIMARY KEY, -- Automatically generates an ID (auto-increment)
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

INSERT INTO users (name, email, phone, password)
VALUES 
    ('John Doe', 'john.doe@example.com', '123-456-7890', 'hashed_password1'),
    ('Jane Smith', 'jane.smith@example.com', '098-765-4321', 'hashed_password2'),
    ('Alice Johnson', 'alice.johnson@example.com', '456-123-7890', 'hashed_password3');


CREATE TABLE SYS_LOG (
    id SERIAL PRIMARY KEY,
    action_datetime TIMESTAMP,
    path_name VARCHAR(255),
    method VARCHAR(50),
    ip VARCHAR(50),
    status_response INT,
    response TEXT,
    description TEXT,
    request_body TEXT,
    request_query TEXT,
    duration FLOAT
);
