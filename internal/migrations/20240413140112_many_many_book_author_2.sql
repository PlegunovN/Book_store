-- +goose Up
-- +goose StatementBegin
ALTER TABLE test_books.public.book DROP COLUMN author;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
