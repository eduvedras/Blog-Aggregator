-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, apikey)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE apikey = ?;
