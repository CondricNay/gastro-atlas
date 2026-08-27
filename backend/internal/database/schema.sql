CREATE TABLE ingredients (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    slug TEXT UNIQUE NOT NULL,
    description TEXT
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