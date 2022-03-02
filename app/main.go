package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/noritama73/go-readinglist/internal/handler"
)

func main() {
	e := echo.New()

	aos := []string{"*"}
	if os.Getenv("ALLOW_ORIGINS") != "" {
		aos = strings.Split(os.Getenv("ALLOW_ORIGINS"), ",")
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: aos,
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	sqldb := handler.NewSQLService()
	cxthandler := handler.NewItemHandler(sqldb)
	e.GET("/itemList", cxthandler.ListItems)
	e.GET("/item", cxthandler.GetItem)
	e.POST("/item", cxthandler.PutItemData)
	e.PUT("/item", cxthandler.UpdateItemData)
	e.DELETE("/item", cxthandler.DeleteItemData)

	e.Logger.Fatal(e.Start(":8080"))
	defer sqldb.DestructDB()
}
