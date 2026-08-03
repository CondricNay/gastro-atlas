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

-- name: GetIngredientPlaces :many
SELECT
    p.id,
    p.name,
    p.type,
    p.latitude,
    p.longitude,
    ip.relationship,
    ip.start_year,
    ip.end_year,
    ip.notes
FROM ingredient_places ip
JOIN places p ON p.id = ip.place_id
WHERE ip.ingredient_id = $1
ORDER BY ip.start_year;