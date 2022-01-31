-- +goose Up
-- +goose StatementBegin
create table leagues
(
    `alias`            varchar(64) NOT NULL PRIMARY KEY,
    `name`             varchar(64) NOT NULL,
    `strength_weight`  int(11)     NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table leagues;
-- +goose StatementEnd
