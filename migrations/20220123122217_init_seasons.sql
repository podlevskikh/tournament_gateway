-- +goose Up
-- +goose StatementBegin
create table seasons
(
    `alias`       varchar(64) NOT NULL PRIMARY KEY,
    `name`        varchar(64) NOT NULL,
    `date_start`  date        not null,
    `date_finish` date        not null,
    `is_current`  tinyint     NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table seasons;
-- +goose StatementEnd
