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
	e.GET("/item", dbhandler.GetItem)
	e.POST("/item", dbhandler.PutItemData)
	e.PUT("/item", dbhandler.UpdateItemData)
	e.DELETE("/item", dbhandler.DeleteItemData)

	e.Logger.Fatal(e.Start(":8080"))
	defer sqldb.DestructDB()
}
