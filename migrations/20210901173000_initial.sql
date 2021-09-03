-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS songs
(
    id     bigserial primary key,
    author text    not null,
    name   text    not null,
    year   integer not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS songs;
-- +goose StatementEnd
