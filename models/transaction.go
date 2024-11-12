package models

import (
	"time"
	//"golang.org/x/text/date"
)

type Transaction struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	BookID      int       `json:"book_id"`
	Quantity    int       `json:"quantity"`
	FinalAmount float64   `json:"final_amount"`
	Status      string    `json:"status"`
	Date        time.Time `json:"date"`
}
