-- name: CreateConfigParam :one
INSERT INTO config (id, created_at, updated_at, name, value)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;


-- name: GetConfigParams :many
SELECT * FROM config;


-- name: UpdateConfigParam :one
UPDATE config SET
updated_at = $2,
value = $3
WHERE name = $1
RETURNING *;



-- name: DeleteConfigParam :exec
DELETE FROM config
WHERE name = $1;