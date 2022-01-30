-- +goose Up
-- +goose StatementBegin
create table teams
(
    `id`          int(11)      NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `name`        varchar(256) NOT NULL,
    `description` text         NOT NULL,
    `foundation`  varchar(256) NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table teams;
-- +goose StatementEnd
