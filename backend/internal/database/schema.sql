CREATE TABLE ingredients (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    slug TEXT UNIQUE NOT NULL,
    description TEXT
);

CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    ingredient_id INT NOT NULL REFERENCES ingredients(id),

    description TEXT NOT NULL,
    time_period TEXT NOT NULL,
    location TEXT NOT NULL,

    start_year INT,
    end_year INT,
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION
);