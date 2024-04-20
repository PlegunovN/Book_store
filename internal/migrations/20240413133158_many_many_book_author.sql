-- +goose Up
-- +goose StatementBegin
CREATE TABLE book_author_id (
id bigserial PRIMARY KEY,
author_id bigint references test_books.public.author,
book_id bigint references test_books.public.book
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table book_author_id;

-- +goose StatementEnd
