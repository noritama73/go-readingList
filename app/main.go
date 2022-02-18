package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/noritama73/go-readinglist/handler"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	sqldb := handler.NewSQLService()
	dbhandler := handler.NewItemHandler(sqldb)
	e.GET("/itemList", dbhandler.ListItems)
	e.POST("/item", dbhandler.PutItemData)

	e.Logger.Fatal(e.Start(":8080"))
}
