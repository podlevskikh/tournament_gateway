-- +goose Up
-- +goose StatementBegin
create table `team_group_player`
(
    `team_id`     int(11)      NOT NULL,
    `group_alias` varchar(128) NOT NULL,
    `player_id`   int(11)      NOT NULL,
    index (team_id, group_alias)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table `team_group_player`;
-- +goose StatementEnd
