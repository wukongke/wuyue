package routes

import (
	"net/http"

	"github.com/labstack/echo"

	book "work-codes/wuyue/app/modules/book/controllers"
	demo "work-codes/wuyue/app/modules/demo/controllers"
	home "work-codes/wuyue/app/modules/home/controllers"
	reader "work-codes/wuyue/app/modules/reader/controllers"
)

func Router(app *echo.Echo) {

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})
	r := app.Group("/wuyue/api")
	r.GET("/demo", demo.DemoController.Demo)

	// home
	r.GET("/index/books", home.IndexController.BookList)

	// type
	r.GET("/types", book.TypeController.List)
	r.GET("/types/:id", book.TypeController.Detail)
	r.GET("/categorys", book.TypeController.BookList)

	// book
	r.GET("/books", book.BookController.List)
	r.GET("/books/:bookNo", book.BookController.Detail)
	r.GET("/book", book.BookController.Info)

	//chapter
	r.GET("/chapters", book.ChapterController.List)
	r.GET("/chapters/:chapterNo", book.ChapterController.Detail)

	//reader
	r.GET("/reader", reader.ReaderController.Info)
	r.GET("/titles", reader.ReaderController.Title)
}
