-- +goose Up
-- +goose StatementBegin
create table `group_results`
(
    `id`           int(11)      NOT NULL,
    `group_alias`  varchar(128) NOT NULL,
    `name`         varchar(128) NOT NULL,
    `scoring_type` varchar(128) NOT NULL,
    index (group_alias)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table `group_results`;
-- +goose StatementEnd
