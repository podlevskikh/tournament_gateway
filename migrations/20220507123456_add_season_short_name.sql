-- +goose Up
-- +goose StatementBegin
ALTER TABLE seasons
    ADD short_name varchar(128) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
