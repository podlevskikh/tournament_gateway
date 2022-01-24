-- +goose Up
-- +goose StatementBegin
create table `team2group`
(
    `team_id`         int(11)     NOT NULL,
    `group_alias`     varchar(64) NOT NULL,
    `handicap_wins`   int(1)      NOT NULL,
    `handicap_points` int(1)      NOT NULL,
    index (team_id),
    index (group_alias)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table `team2group`;
-- +goose StatementEnd
