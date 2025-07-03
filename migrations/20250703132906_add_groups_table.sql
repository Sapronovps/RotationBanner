-- +goose Up
-- +goose StatementBegin
CREATE TABLE groups
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(100),
    description VARCHAR(255),
    created_at  TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE groups;
-- +goose StatementEnd
