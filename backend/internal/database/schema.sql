CREATE TABLE ingredients (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    slug TEXT UNIQUE NOT NULL,
    description TEXT
);

CREATE TABLE places (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT NOT NULL
);

CREATE TABLE ingredient_places (
    ingredient_id INT REFERENCES ingredients(id),
    place_id INT REFERENCES places(id),

    relationship TEXT NOT NULL,

    start_year INT,
    end_year INT,

    notes TEXT,

    PRIMARY KEY (
        ingredient_id,
        place_id
    )
);