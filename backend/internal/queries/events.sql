-- name: CreateEvent :one
INSERT INTO events (
    ingredient_id,
    title,
    description,
    time_period,
    entity,
    location,
    sources,
    confidence
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: GetEventsByIngredient :many
SELECT *
FROM events
WHERE ingredient_id = $1
ORDER BY id;