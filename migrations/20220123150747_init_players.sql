-- +goose Up
-- +goose StatementBegin
create table players
(
    `id`          int(11)     NOT NULL AUTO_INCREMENT,
    `first_name`  varchar(64) NOT NULL,
    `last_name`   varchar(64) NOT NULL,
    `middle_name` varchar(64) NOT NULL,
    `birth_date`  date         DEFAULT NULL,
    `height`      int(4)       DEFAULT NULL,
    `gender`      varchar(16) NOT NULL,
    `skill`       varchar(64)  DEFAULT '',
    `avatar_url`  varchar(256) DEFAULT '',
    PRIMARY KEY (`id`)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table players;
-- +goose StatementEnd
