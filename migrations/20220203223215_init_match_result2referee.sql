-- +goose Up
-- +goose StatementBegin
create table `match_result2referee`
(
    `match_result_match_id` int(11) NOT NULL,
    `referee_id`            int(11) NOT NULL,
    unique index (match_result_match_id, referee_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table `match_result2referee`;
-- +goose StatementEnd
