-- +goose Up
-- +goose StatementBegin
create table referees
(
    `id`          int(11)     NOT NULL AUTO_INCREMENT,
    `first_name`  varchar(64) NOT NULL,
    `last_name`   varchar(64) NOT NULL,
    `middle_name` varchar(64) NOT NULL,
    `phone`       varchar(64)  DEFAULT '',
    `skill`       varchar(256) DEFAULT '',
    `description` text         DEFAULT NULL,
    `price`       varchar(256) DEFAULT '',
    PRIMARY KEY (`id`)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table referees;
-- +goose StatementEnd
