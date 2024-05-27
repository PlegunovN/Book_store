-- +goose Up
-- +goose StatementBegin
ALTER TABLE test_books.public.book_author_id RENAME TO  bookAuthor
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
