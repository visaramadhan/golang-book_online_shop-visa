package models

type Book struct {
	ID			int		`json:"id"`
	Name		string	`json:"name"`
	Type 		string	`json:"type"`
	Pengarang	string	`json:"pengarang"`
	Harga		string	`json:"harga"`
}