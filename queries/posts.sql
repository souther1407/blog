-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, content, author_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id,created_at, updated_at, title;



-- name: GetLastPosts :many
SELECT p.id,p.title, p.created_at, a.name as author FROM posts p
JOIN users a on a.id = p.author_id
LIMIT $1;



-- name: UpdatePost :one
UPDATE posts SET
updated_at = $2,
title = $3,
content = $4 
WHERE id = $1
RETURNING *;



-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1 AND author_id = $2;