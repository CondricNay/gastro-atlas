CREATE TABLE ingredients (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    slug TEXT UNIQUE NOT NULL,
    description TEXT
);

CREATE TABLE places (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    type TEXT NOT NULL,
    latitude  DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL
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

CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    ingredient_id INT NOT NULL REFERENCES ingredients(id),

    title TEXT NOT NULL,
    description TEXT NOT NULL,
    time_period TEXT NOT NULL,
    entity TEXT NOT NULL,
    location TEXT NOT NULL,

    sources TEXT[] NOT NULL,
    confidence TEXT NOT NULL
);