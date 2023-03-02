package main

import (
	"echo/database"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	database.DBConnection()

	e := echo.New()

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	log.Fatal(e.Start(":3000"))
}
