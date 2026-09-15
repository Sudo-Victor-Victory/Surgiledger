-- name: CreateEvent :one
INSERT INTO events (
    episode_id,
    event_type,
    payload
)
VALUES ($1, $2, $3)
RETURNING *;


-- name: GetEventById :one
SELECT *
FROM events
WHERE episode_id = $1
AND id = $2;

-- name: GetEpisodeEvents :many
SELECT *
FROM events
WHERE episode_id = $1
ORDER BY created_at ASC;