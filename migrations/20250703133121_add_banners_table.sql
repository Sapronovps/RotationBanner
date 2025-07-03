-- +goose Up
-- +goose StatementBegin
CREATE TABLE banners
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(100),
    description VARCHAR(255),
    created_at  TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE banners;
-- +goose StatementEnd
