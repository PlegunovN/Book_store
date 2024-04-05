package books

type Book struct {
	ID       int64  `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	AuthorID int64  `json:"author" db:"author"`
}

type Author struct {
	ID        int64  `json:"id" db:"id"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname  string `json:"lastname" db:"lastname"`
}
type BookAuthor struct {
	ID        int64  `json:"id" db:"id"`
	Title     string `json:"title" db:"title"`
	AuthorID  int64  `json:"author" db:"author"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname  string `json:"lastname" db:"lastname"`
}
