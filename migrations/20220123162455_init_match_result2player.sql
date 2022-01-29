-- +goose Up
-- +goose StatementBegin
create table `match_result2player`
(
    `match_result_match_id` int(11)     NOT NULL,
    `player_id`       int(11)     NOT NULL,
    `team`            varchar(16) NOT NULL,
    index (match_result_match_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table `match_result2player`;
-- +goose StatementEnd
