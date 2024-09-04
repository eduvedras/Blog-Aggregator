-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds ORDER BY created_at ASC LIMIT ? OFFSET ?;

-- name: MarkFeedFetch :one
UPDATE feeds
SET last_fetched_at = CURRENT_TIMESTAMP,
updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT ?;
