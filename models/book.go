package models

// Model Book
type Book struct {
	ID     uint64 `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json: "author"`
}
