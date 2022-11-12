package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var books []Book

func main() {
	e := echo.New()
	// CORSの設定
	e.Use(middleware.CORS())

	e.GET("/api/books", getBooks)
	// e.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	// e.HandleFunc("/api/books", createBook).Methods("POST")
	// e.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	// e.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	port := os.Getenv("API_PORT")

	e.Logger.Fatal(e.Start(":" + port))
}

func getBooks(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
