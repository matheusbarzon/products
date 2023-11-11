package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {

	var server string = os.Getenv("MYSQL_SERVER")
	var database string = os.Getenv("MYSQL_DATABASE")
	var user string = os.Getenv("MYSQL_USER")
	var pass string = os.Getenv("MYSQL_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, server, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
