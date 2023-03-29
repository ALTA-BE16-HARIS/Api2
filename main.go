package main

import (
	"belajarAPI/books"
	"belajarAPI/config"
	"belajarAPI/users"

	"github.com/labstack/echo"
)

func main() {
	_, _ = config.InitDB()
	e := echo.New()
	e.POST("/login", users.Login())
	e.POST("/register", users.Register())
	e.POST("/books", books.Addbook())
	e.GET("/books", books.Getallbooks())
	e.Start(":3758")
}
