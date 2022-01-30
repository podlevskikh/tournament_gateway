-- +goose Up
-- +goose StatementBegin
create table gym2metro
(
    `gym_id`   int(11),
    `metro_id` int(11)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
