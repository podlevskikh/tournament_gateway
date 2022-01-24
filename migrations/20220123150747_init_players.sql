-- +goose Up
-- +goose StatementBegin
create table players
(
    `id`          int(11)      NOT NULL PRIMARY KEY,
    `first_name`  varchar(64)  NOT NULL,
    `last_name`   varchar(64)  NOT NULL,
    `middle_name` varchar(64)  NOT NULL,
    `birth_date`  date  NOT NULL,
    `gender`      varchar(16)  NOT NULL,
    `skill`       varchar(64)  NOT NULL,
    `avatar_url`  varchar(256) NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table players;
-- +goose StatementEnd
