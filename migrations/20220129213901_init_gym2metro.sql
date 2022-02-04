-- +goose Up
-- +goose StatementBegin
create table gym2metro
(
    `gym_id`   int(11),
    `metro_id` int(11),
    UNIQUE KEY `gym_id` (`gym_id`, `metro_id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table gym2metro;
-- +goose StatementEnd
