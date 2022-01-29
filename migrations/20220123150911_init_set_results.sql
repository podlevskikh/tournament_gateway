-- +goose Up
-- +goose StatementBegin
create table `set_results`
(
    `result_match_id`    int(11) NOT NULL,
    `set_number`  int(11) NOT NULL,
    `home_score`  int(11) NOT NULL,
    `guest_score` int(11) NOT NULL,
    index (result_match_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table `set_results`;
-- +goose StatementEnd
