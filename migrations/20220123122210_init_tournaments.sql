-- +goose Up
-- +goose StatementBegin
create table tournaments
(
    `alias`       varchar(64)  NOT NULL PRIMARY KEY,
    `name`        varchar(64)  NOT NULL,
    `description` varchar(512) NOT NULL,
    `gender`      varchar(64)  NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table tournaments;
-- +goose StatementEnd
