-- +goose Up
-- +goose StatementBegin
create table `match_results`
(
    `match_id`                 int(11)    NOT NULL PRIMARY KEY,
    `match_datetime`           datetime,

    `home_points`              int(11)    NOT NULL,
    `home_best_player_id`      int(11),
    `home_referee_evaluation`  int(2),

    `guest_points`             int(11)    NOT NULL,
    `guest_best_player_id`     int(11)    NOT NULL,
    `guest_referee_evaluation` int(2)     NOT NULL,

    `winner`                   varchar(8) NOT NULL,
    `referee_id`               int(11)    NOT NULL,

    `approved`                 tinyint    NOT NULL DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table `match_results`;
-- +goose StatementEnd
