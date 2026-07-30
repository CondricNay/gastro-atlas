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


-- name: CreateIngredient :one
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

RETURNING
    id,
    name,
    slug,
    description;