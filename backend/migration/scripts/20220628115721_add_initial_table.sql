-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS athletes(
    id TEXT PRIMARY KEY NOT NULL,
    first_name VARCHAR(20) NOT NULL, 
    last_name VARCHAR(30) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS workouts(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    description VARCHAR(200),
    athlete_id TEXT NOT NULL REFERENCES athletes (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS athletes;
DROP TABLE IF EXISTS workouts;
-- +goose StatementEnd
