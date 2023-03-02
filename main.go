package main

import (
	"echo/database"
	"echo/database/entity"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	database.DBConnection()

	e := echo.New()

	// hello world 예시
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	// route
	e.GET("/users", fetchUsers)
	e.GET("/users/:id", getUser)

	// 직접 해보아요
	e.POST("/users/:id", createUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	log.Fatal(e.Start(":3000"))
}

func fetchUsers(c echo.Context) error {
	db := database.DBConnection()
	var users []entity.Users

	// 모든 유저의 레코드
	// 참고: https://gorm.io/docs/query.html
	db.Find(&users)

	return c.JSON(http.StatusOK, users)
}

func getUser(c echo.Context) error {
	db := database.DBConnection()
	var users entity.Users

	// path parameters
	// 참고: https://echo.labstack.com/guide/request/#path-parameters
	id := c.Param("id")

	db.Where("id = ?", id).Find(&users)

	return c.JSON(http.StatusOK, users)
}

// 아래를 채워주세요
// 공식문서 참고
// echo: https://echo.labstack.com/guide/
// gorm: https://gorm.io/ko_KR/docs/index.html
func createUser(c echo.Context) error {
	return nil
}

func updateUser(c echo.Context) error {
	return nil
}

func deleteUser(c echo.Context) error {
	return nil
}
