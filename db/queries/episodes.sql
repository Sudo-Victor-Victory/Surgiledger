-- name: CreateEpisode :one
INSERT INTO episodes (status)
VALUES ($1)
RETURNING *;


-- name: GetEpisode :one
SELECT *
FROM episodes
WHERE id = $1;


-- name: ListEpisodes :many
SELECT *
FROM episodes
ORDER BY created_at DESC;


-- name: UpdateEpisodeStatus :one
UPDATE episodes
SET
    status = ($2),
    updated_at = NOW()
WHERE id = ($1)
RETURNING *;


-- name: DeleteEpisode :exec
DELETE FROM
episodes WHERE id = $1;


