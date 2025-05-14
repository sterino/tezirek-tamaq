-- +goose Up
CREATE TABLE users (
                       id UUID PRIMARY KEY,
                       name TEXT NOT NULL,
                       email TEXT NOT NULL UNIQUE,
                       password TEXT NOT NULL,
                       role TEXT NOT NULL,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE users;
