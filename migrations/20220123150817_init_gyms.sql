-- +goose Up
-- +goose StatementBegin
create table gyms
(
    `id`            int(11) NOT NULL PRIMARY KEY,
    `metro_station` text    NOT NULL,
    `address`       text    NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table gyms;
-- +goose StatementEnd
