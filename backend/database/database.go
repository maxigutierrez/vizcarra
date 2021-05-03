package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func InitDb() {
	// connect to database
	conn, err := gorm.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/vizcarra?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	db = conn
	db.LogMode(true)
	fmt.Println("conn")
}

func GetDb() *gorm.DB {
	return db
}
