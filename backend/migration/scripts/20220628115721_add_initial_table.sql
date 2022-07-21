-- +goose Up
-- +goose StatementBegin
-- CREATE TABLE IF NOT EXISTS athletes(
--     id TEXT PRIMARY KEY NOT NULL,
--     first_name VARCHAR(20) NOT NULL, 
--     last_name VARCHAR(30) NOT NULL,
--     created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
-- );

CREATE TABLE IF NOT EXISTS workouts(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    description VARCHAR(200)
);

CREATE TABLE IF NOT EXISTS excercises (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    workout_id INT NOT NULL REFERENCES workouts (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS sets (
    id SERIAL PRIMARY KEY,
    weight FLOAT NOT NULL,
    reps INT NOT NULL,
    excercise_id INT NOT NULL REFERENCES excercises (id) ON DELETE CASCADE
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sets;
DROP TABLE IF EXISTS excercises;
DROP TABLE IF EXISTS workouts;
-- +goose StatementEnd
