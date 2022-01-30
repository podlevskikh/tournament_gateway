-- +goose Up
-- +goose StatementBegin
create table `matches`
(
    `id`            int(11)     NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `date`          date,
    `home_team_id`  int(11)     NOT NULL,
    `guest_team_id` int(11)     NOT NULL,
    `stage_alias`   varchar(64) NOT NULL,
    `league_alias`  varchar(64) NOT NULL,
    `group_alias`   varchar(64) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table `matches`;
-- +goose StatementEnd
