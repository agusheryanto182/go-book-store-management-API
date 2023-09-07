package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPassword := "root"
	dbName := "book"

	d, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		fmt.Println("connect to database failed")
		return
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
