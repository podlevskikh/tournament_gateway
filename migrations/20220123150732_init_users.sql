-- +goose Up
-- +goose StatementBegin
create table users
(
    `id`          int(11)      NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `first_name`  varchar(256)  NOT NULL,
    `last_name`   varchar(256)  NOT NULL,
    `middle_name` varchar(256)  NOT NULL,
    `nick_name`   varchar(256) NOT NULL,
    `email`       varchar(256) NOT NULL,
    `phone`       varchar(256)  NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
