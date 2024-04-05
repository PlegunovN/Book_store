-- +goose Up
-- +goose StatementBegin
create table "user"
(
    id bigserial primary key,
    name varchar(50),
    email varchar(50),
    password varchar(256),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table "user";
-- +goose StatementEnd
