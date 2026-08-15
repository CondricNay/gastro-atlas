-- name: GetIngredients :many
SELECT
    id,
    name,
    slug,
    description
FROM ingredients
ORDER BY name;


-- name: GetIngredientBySlug :one
SELECT
    id,
    name,
    slug,
    description
FROM ingredients
WHERE slug = $1;


-- name: UpsertIngredient :one
INSERT INTO ingredients (
    name,
    slug,
    description
)
VALUES (
    $1,
    $2,
    $3
)
ON CONFLICT (slug)
DO UPDATE SET
    name = EXCLUDED.name,
    description = EXCLUDED.description
RETURNING id;