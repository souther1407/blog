-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, content, author_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id,created_at, updated_at, title;