package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	api := e.Group("/api")
	api.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
