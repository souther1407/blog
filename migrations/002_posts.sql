-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT UNIQUE NOT NULL,
    content TEXT NOT NULL,
    author_id UUID NOT NULL REFERENCES users(id)
);
-- +goose Down
DROP TABLE posts;