package repository

import (
	"book_online_shop_visa/models"
	"database/sql"
	"fmt"
)

type BookRepository interface {
    CreateModel(book *models.Book) error
	GetModelsByID(id int)(*models.Book, error)
}

type bookRepository struct {
    db *sql.DB
}

// NewBookRepository create new bookrepository
func NewBookRepository(db *sql.DB) bookRepository{
	return bookRepository{db: db}
}

// CreateModel insert new book into database
func (r *bookRepository) CreateBook (book *models.Book){
	sqlStatement := `INSERT INTO books (title, author, price) VALUES ($1, $2, $3)`
    _, err := r.db.Exec(sqlStatement, book.Name, book.Pengarang, book.Harga)
    if err!= nil {
        panic(err.Error())
    }
    fmt.Println("Buku Berhasil Ditambahkan")
}

// GetModelsByID get book by id
func (r *bookRepository) GetBookbyID (id int) (*models.Book){
	var book models.Book
    row := r.db.QueryRow("SELECT id, name, author, price FROM books WHERE id=$1", id)
    err := row.Scan(&book.ID, &book.Name, &book.Pengarang, &book.Harga)
    if err == sql.ErrNoRows {
        return nil
    } else if err!= nil {
        panic(err.Error())
    }
    return &book
}