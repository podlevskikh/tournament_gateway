-- +goose Up
-- +goose StatementBegin
create table `team2user`
(
    `team_id`   int(11) NOT NULL,
    `user_id`  int(11) NOT NULL,
    index (team_id, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table `team2user`;
-- +goose StatementEnd
