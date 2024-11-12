package main

import(
	"database/sql"
	"fmt"
	"log"
	// "github.com/lib/pq"
)

var db *sql.DB

func connectDB() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=visa281101 dbname=book_online_shop sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err!= nil {
        return nil, err
    }


	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Berhasil terhubung ke database!")
	return db, nil
}

func init() {
	var err error
    db, err = connectDB()
    if err!= nil {
        log.Fatal(err)
    }
}