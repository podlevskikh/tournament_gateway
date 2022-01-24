-- +goose Up
-- +goose StatementBegin
create table stages
(
    `alias`        varchar(64)  NOT NULL PRIMARY KEY,
    `name`         varchar(64)  NOT NULL,
    `date_start`   date         NOT NULL,
    `date_finish`  date         NOT NULL,
    `is_current`   tinyint      NOT NULL,
    `icon_url`     varchar(256) NOT NULL,
    `season_alias` varchar(64)  NOT NULL,
    `type`         varchar(64)  NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table stages;
-- +goose StatementEnd
