-- +goose Up
-- +goose StatementBegin
create table metros
(
    `id`   int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `name` text    NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table metros;
-- +goose StatementEnd
