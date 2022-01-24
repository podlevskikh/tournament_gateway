-- +goose Up
-- +goose StatementBegin
create table teams
(
    `id`              int(11)      NOT NULL PRIMARY KEY,
    `name`            varchar(256) NOT NULL,
    `description`     text         NOT NULL,
    `foundation_date` date         NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table teams;
-- +goose StatementEnd
