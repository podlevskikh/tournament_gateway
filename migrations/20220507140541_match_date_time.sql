-- +goose Up
-- +goose StatementBegin
ALTER TABLE matches
    MODIFY `date` datetime;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
