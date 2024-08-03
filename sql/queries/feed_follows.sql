-- name: CreateFeedFollow :one
INSERT INTO feed_follows(id, created_at, updated_at, feed_id, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetAllFeedFollowsByUserId :many
SELECT * FROM feed_follows 
WHERE user_id = $1
ORDER BY created_at;

-- name: DeleteFeedFollowById :one
DELETE FROM feed_follows WHERE id = $1 and user_id = $2
RETURNING *;
