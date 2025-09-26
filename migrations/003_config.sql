-- +goose Up
CREATE TABLE config (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL UNIQUE,
    value TEXT NOT NULL
);


-- +goose Down
DROP TABLE config;