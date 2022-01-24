-- +goose Up
-- +goose StatementBegin
create table `team2gym`
(
    `team_id`      int(11)    NOT NULL,
    `gym_id`       int(11)    NOT NULL,
    `week_day`     int(1)     NOT NULL,
    `time_from`    varchar(8) NOT NULL,
    `time_to`      varchar(8) NOT NULL,
    `time_warm_up` varchar(8) NOT NULL,
    index (team_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table `team2gym`;
-- +goose StatementEnd
