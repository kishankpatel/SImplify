package services

import (
	"database/sql"
	"fmt"

	"github.com/kishankpatel/simplify/models"
)

// BookService - to take actions on book
type BookService interface {
	// AddNewBook() models.Book
	// Books() ([]*models.Book, error)
	AddNewBook(book *models.Book, db *sql.DB) models.Book
	AllBooks(db *sql.DB) []models.Book
}

type bookService struct {
	Creater string
}

// NewBookService - constructure to create a new book
func NewBookService() (BookService, error) {
	var bookService BookService = &bookService{"Kishan Patel"}
	return bookService, nil
}

func (b bookService) AddNewBook(book *models.Book, db *sql.DB) models.Book {
	// stmt, err := db.Prepare("INSERT INTO books(name, author, price, created_by) VALUES(?,?,?,?)")
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	// bookData, err := stmt.Exec(book.Name, book.Author, book.Price, b.Creater)
	// var id int
	sql := `
    insert into books (name, author, price, created_by)
    values ($1, $2, $3, $4) RETURNING id;
	`
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	res, err := stmt.Exec(book.Name, book.Author, book.Price, b.Creater)
	// err = db.QueryRow("insert into books (name, author, price, created_by) VALUES ($1, $2, $3, $4)", book.Name, book.Author, book.Price, b.Creater).Scan(&id)
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("id====>", id)
	selDB, err := db.Query("SELECT * FROM books where id =($1)", string(id))
	if err != nil {
		panic(err.Error())
	}
	bookRes := models.Book{}
	for selDB.Next() {
		var id, name, author, price, created_by, created_at string
		err = selDB.Scan(&id, &name, &author, &price, &created_by, &created_at)
		if err != nil {
			panic(err.Error())
		}
		bookRes.ID = id
		bookRes.Name = name
		bookRes.Author = author
		bookRes.Price = price
		bookRes.CreatedBy = created_by
	}
	return bookRes
}

func (b bookService) AllBooks(db *sql.DB) []models.Book {
	selDB, err := db.Query("SELECT * FROM books ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	book := models.Book{}
	res := []models.Book{}
	for selDB.Next() {
		var id, name, author, price, created_by, created_at string
		err = selDB.Scan(&id, &name, &author, &price, &created_by, &created_at)
		if err != nil {
			panic(err.Error())
		}
		book.ID = id
		book.Name = name
		book.Author = author
		book.Price = price
		book.CreatedBy = created_by
		book.CreatedAt = created_at
		res = append(res, book)
	}
	return res
}
