-- +goose Up
-- +goose StatementBegin
CREATE TABLE slots
(
    id          SERIAL PRIMARY KEY,
    description VARCHAR(255),
    created_at  TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE slots;
-- +goose StatementEnd
