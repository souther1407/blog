-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at, name, email, password)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;


-- name: GetUser :one
SELECT * FROM users WHERE name = $1;