-- +goose Up
-- +goose StatementBegin
create table gyms
(
    `id`          int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `address`     text    NOT NULL,
    `name`        text DEFAULT NULL,
    `description` text DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table gyms;
-- +goose StatementEnd
