-- +goose Up
ALTER TABLE posts ADD description TEXT;

-- +goose Down
ALTER TABLE posts DROP COLUMN description;