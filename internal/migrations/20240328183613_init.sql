-- +goose Up
-- +goose StatementBegin
create table author
(
    id        bigserial
        primary key,
    firstname varchar(50),
    lastname  varchar(50)
);
create table book
(
    id     bigserial
        primary key,
    title  varchar(50),
    author bigint
        references author
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table book;
drop table author;

-- +goose StatementEnd
