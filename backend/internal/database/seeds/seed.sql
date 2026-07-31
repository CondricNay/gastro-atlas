-- Ingredients

INSERT INTO ingredients (
    slug,
    name,
    description
)
VALUES (
    'tomato',
    'Tomato',
    'A fruit native to the Americas that became a major ingredient in cuisines worldwide.'
)
ON CONFLICT (slug) DO NOTHING;

-- Places

INSERT INTO places (
    name,
    type
)
VALUES
(
    'Andes Region',
    'region'
),
(
    'Mesoamerica',
    'region'
),
(
    'Spanish Empire',
    'empire'
),
(
    'Italy',
    'country'
)
ON CONFLICT DO NOTHING;


-- Ingredient Place Relationships
INSERT INTO ingredient_places (
    ingredient_id,
    place_id,
    relationship,
    start_year,
    end_year,
    notes
)
VALUES
(
    (SELECT id FROM ingredients WHERE slug = 'tomato'),
    (SELECT id FROM places WHERE name = 'Andes Region'),
    'origin',
    -500,
    1500,
    'Wild tomato relatives and early cultivation are associated with the Andes region.'
),
(
    (SELECT id FROM ingredients WHERE slug = 'tomato'),
    (SELECT id FROM places WHERE name = 'Mesoamerica'),
    'cultivation',
    -500,
    1500,
    'Tomatoes were cultivated and incorporated into local cuisines.'
),
(
    (SELECT id FROM ingredients WHERE slug = 'tomato'),
    (SELECT id FROM places WHERE name = 'Spanish Empire'),
    'introduced',
    1500,
    NULL,
    'Tomatoes were brought to Europe after Spanish contact with the Americas.'
),
(
    (SELECT id FROM ingredients WHERE slug = 'tomato'),
    (SELECT id FROM places WHERE name = 'Italy'),
    'popularized',
    1700,
    NULL,
    'Tomatoes became central to Italian cuisine, especially sauces.'
)
ON CONFLICT DO NOTHING;