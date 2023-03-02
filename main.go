package main

import (
	"echo/database"
	"echo/database/entity"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

func main() {
	database.DBConnection()

	e := echo.New()
	e.Use(middleware.Logger())

	// hello world 예시
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	// route
	e.GET("/users", fetchUsers)
	e.GET("/users/:id", getUser)

	log.Fatal(e.Start(":3000"))
}

func fetchUsers(c echo.Context) error {
	db := database.DBConnection()

	var users []entity.Users
	db.Find(&users)

	return c.JSON(http.StatusOK, users)
}

func getUser(c echo.Context) error {
	db := database.DBConnection()
	var users entity.Users
	var binder struct {
		Id int64 `json:"-" xml:"-" param:"id"`
	}
	
	if err := c.Bind(&binder); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}

	db.Where("id = ?", binder.Id).Find(&users)

	return c.JSON(http.StatusOK, users)
}
