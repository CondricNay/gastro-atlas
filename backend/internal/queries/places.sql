-- name: UpsertPlace :one
INSERT INTO places (
    name,
    type,
    latitude,
    longitude
)
VALUES (
    $1,
    $2,
    $3,
    $4
)
ON CONFLICT (name)
DO UPDATE SET
    type = EXCLUDED.type,
    latitude = EXCLUDED.latitude,
    longitude = EXCLUDED.longitude
RETURNING id;