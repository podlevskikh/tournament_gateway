-- +goose Up
-- +goose StatementBegin
create table seasons
(
    `alias`       varchar(64) NOT NULL,
    `name`        varchar(64) NOT NULL,
    `date_start`  date        NOT NULL,
    `date_finish` date        NOT NULL,
    `is_current`  tinyint(4)  NOT NULL,
    PRIMARY KEY (`alias`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table seasons;
-- +goose StatementEnd
