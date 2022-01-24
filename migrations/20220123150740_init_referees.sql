-- +goose Up
-- +goose StatementBegin
create table referees
(
    `id`          int(11)      NOT NULL PRIMARY KEY,
    `first_name`  varchar(64)  NOT NULL,
    `last_name`   varchar(64)  NOT NULL,
    `middle_name` varchar(64)  NOT NULL,
    `phone`       varchar(64)  NOT NULL,
    `skill`       varchar(256) NOT NULL,
    `description` text         NOT NULL,
    `price`       varchar(256) NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table referees;
-- +goose StatementEnd
