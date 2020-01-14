package models

// Book struct declaration
type Book struct {
	ID        string
	Name      string `json:"name"`
	Author    string `json:"author"`
	Price     string `json:"price"`
	CreatedBy string
	CreatedAt string
}

// NewBook method declaration
func NewBook(book *Book) *Book {

	return &Book{
		ID:     "",
		Name:   book.Name,
		Author: book.Author,
		Price:  book.Price,
	}
}
