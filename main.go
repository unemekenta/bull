package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/unemekenta/bull/docs"
)

type Response struct {
	Int64  int64  `json:"int64"`
	String string `json:"string"`
	World  *Item  `json:"world"`
}

type Item struct {
	Text string `json:"text"`
}

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

// @title Swagger Example API
// @version 1.0
// @description This is a sample swagger server.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:1314
// @BasePath /api/v1
func main() {
	e := echo.New()
	// CORSの設定
	e.Use(middleware.CORS())

	initRouting(e)

	port := os.Getenv("API_PORT")

	e.Logger.Fatal(e.Start(":" + port))
}

func initRouting(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/api/v1/books", getBooks)

	// e.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	// e.HandleFunc("/api/books", createBook).Methods("POST")
	// e.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	// e.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

}

// hello godoc
// @Summary  Hello World !
// @ID       HelloWorldIndex
// @Tags     HelloWorld
// @Produce  json
// @Success  200  {object}  Response
// @Router   / [get]
func getBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, &Response{
		Int64:  1,
		String: "example",
		World: &Item{
			Text: "hello world !",
		},
	})
}
