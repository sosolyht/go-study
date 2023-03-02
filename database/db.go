package database

import (
	"echo/database/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var db *gorm.DB

func DBConnection() (db *gorm.DB) {
	// gorm 의 mysql 드라이버를 이용해 기존 dsn 으로 연결
	// Config 의 Logger 를 통한 쿼리 로그
	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/study"),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&entity.Users{},
	)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetDBInstance() *gorm.DB {
	return db
}
