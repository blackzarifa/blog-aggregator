-- name: CreateFeedFollow :one
WITH inserted_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES ($1, $2, $3, $4, $5)
  RETURNING *
)
SELECT 
  f.*,
  u.name AS user_name,
  fd.name AS feed_name
FROM inserted_follow f
JOIN users u ON f.user_id = u.id
JOIN feeds fd ON f.feed_id = fd.id;

-- name: GetFeedFollowsForUser :many
SELECT 
  ff.*,
  f.name AS feed_name,
  u.name AS user_name
FROM feed_follows ff
JOIN feeds f ON ff.feed_id = f.id
JOIN users u ON ff.user_id = u.id
WHERE ff.user_id = $1
ORDER BY f.name ASC;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows ff
WHERE ff.user_id = $1 AND ff.feed_id = (
  SELECT id FROM feeds WHERE url = $2
);

