-- +goose Up
-- +goose StatementBegin
create table `groups`
(
    `alias`            varchar(128) NOT NULL PRIMARY KEY,
    `short_name`       varchar(32)  NOT NULL,
    `name`             varchar(256) NOT NULL,
    `description`      varchar(256) NOT NULL,
    `tournament_alias` varchar(64)  NOT NULL,
    `season_alias`     varchar(64)  NOT NULL,
    `stage_alias`      varchar(64)  NOT NULL,
    `league_alias`     varchar(64)  NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table `groups`;
-- +goose StatementEnd
