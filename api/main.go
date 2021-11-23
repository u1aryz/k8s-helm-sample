package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type result struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	api := e.Group("/api")
	api.GET("/your-name", func(c echo.Context) error {
		val, err := client.Get("name").Result()

		if err == redis.Nil {
			return echo.NewHTTPError(http.StatusNotFound, "What's your name?")
		}
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, &result{
			Message: fmt.Sprintf("Hello %s", val),
		})
	})

	api.POST("/your-name", func(c echo.Context) error {
		value := c.FormValue("name")
		if value == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "You must 'name' form data")
		}

		err := client.Set("name", value, 1*time.Minute).Err()

		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, &result{
			Message: fmt.Sprintf("Hello %s", value),
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
