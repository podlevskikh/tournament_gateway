-- +goose Up
-- +goose StatementBegin
create table `team_group_player`
(
    `team_id`     int(11)      NOT NULL,
    `group_alias` varchar(128) NOT NULL,
    `player_id`   int(11)      NOT NULL,
    UNIQUE KEY `team_group_player` (`team_id`, `group_alias`, `player_id`)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table `team_group_player`;
-- +goose StatementEnd
