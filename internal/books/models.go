package books

type Book struct {
	ID        int64  `json:"id" db:"id"`
	BookTitle string `json:"title" db:"title"`
}

type Author struct {
	ID              int64  `json:"id" db:"id"`
	AuthorFirstname string `json:"firstname" db:"firstname"`
	AuthorLastname  string `json:"lastname" db:"lastname"`
}

type BookAuthor struct {
	ID              int64  `json:"id" db:"id"`
	Title           string `json:"title" db:"title"`
	AuthorID        int64  `json:"author" db:"author"`
	AuthorFirstname string `json:"firstname" db:"firstname"`
	AuthorLastname  string `json:"lastname" db:"lastname"`
}
