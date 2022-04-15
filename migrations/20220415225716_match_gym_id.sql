-- +goose Up
-- +goose StatementBegin
ALTER TABLE matches
    ADD gym_id int(11) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
