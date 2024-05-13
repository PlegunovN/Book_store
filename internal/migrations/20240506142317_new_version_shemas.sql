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
    title  varchar(50)
    );
create table "user"
(
    id bigserial primary key,
    name varchar(50),
    email varchar(50),
    password varchar(256),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);
CREATE TABLE bookAuthor(
    id bigserial PRIMARY KEY,
    author_id bigint references author,
    book_id bigint references book
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table book;
drop table author;
drop table "user";
drop table bookAuthor;
-- +goose StatementEnd
