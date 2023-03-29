package users

import (
	"belajarAPI/books"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	book     []books.Book
}
