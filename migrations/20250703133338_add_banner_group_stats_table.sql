-- +goose Up
-- +goose StatementBegin
CREATE TABLE banner_group_stats
(
    id         SERIAL PRIMARY KEY,
    slot_id    INTEGER,
    banner_id  INTEGER,
    group_id   INTEGER,
    clicks     INTEGER,
    shows      INTEGER,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

ALTER TABLE banner_group_stats
    ADD CONSTRAINT fk_slot_id FOREIGN KEY (slot_id) REFERENCES slots (id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE banner_group_stats
    ADD CONSTRAINT fk_banner_id FOREIGN KEY (banner_id) REFERENCES banners (id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE banner_group_stats
    ADD CONSTRAINT fk_group_id FOREIGN KEY (group_id) REFERENCES groups (id) ON UPDATE CASCADE ON DELETE CASCADE;

CREATE INDEX idx_banner_group_stats_slot_id_group_id ON banner_group_stats (slot_id, group_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE banner_group_stats;
-- +goose StatementEnd
