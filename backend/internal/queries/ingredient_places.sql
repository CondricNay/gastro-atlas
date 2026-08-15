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


-- name: UpsertIngredientPlace :exec
INSERT INTO ingredient_places (
    ingredient_id,
    place_id,
    relationship,
    start_year,
    end_year,
    notes
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
ON CONFLICT (ingredient_id, place_id)
DO UPDATE SET
    relationship = EXCLUDED.relationship,
    start_year = EXCLUDED.start_year,
    end_year = EXCLUDED.end_year,
    notes = EXCLUDED.notes;