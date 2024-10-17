-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByName :one
SELECT * FROM users
WHERE name = $1 LIMIT 1;

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT * FROM users;

-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.name, feeds.url, users.name AS created_by 
FROM feeds
JOIN users ON feeds.user_id = users.id;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE url = $1 LIMIT 1;

-- name: CreateFeedFollow :one
WITH inserted_feed_follows AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)

SELECT 
    inserted_feed_follows.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follows
INNER JOIN feeds ON inserted_feed_follows.feed_id = feeds.id
INNER JOIN users ON inserted_feed_follows.user_id = users.id;

-- name: GetFeedFollowsForUser :many
SELECT 
    feed_follows.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM feed_follows 
INNER JOIN users ON feed_follows.user_id = users.id
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1; 

-- name: DeleteFollowsForUser :exec
DELETE FROM feed_follows
WHERE feed_follows.user_id = $2 AND (
    SELECT url FROM feeds
    WHERE feeds.id = feed_follows.feed_id
) = $1;