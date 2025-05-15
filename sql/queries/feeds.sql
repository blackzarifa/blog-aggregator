-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;


-- name: GetFeed :one
SELECT *
FROM feeds
WHERE name = $1 LIMIT 1;

-- name: GetFeedByURL :one
SELECT *
FROM feeds
WHERE url = $1
LIMIT 1;


-- name: DeleteAllFeeds :exec
DELETE FROM feeds;


-- name: GetFeeds :many
SELECT * FROM feeds;


-- name: GetFeedsWithUsers :many
SELECT 
  f.id, f.created_at, f.updated_at, f.name, f.url, f.user_id,
  u.name as user_name
FROM feeds f
JOIN users u ON f.user_id = u.id;
