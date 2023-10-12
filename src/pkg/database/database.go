package database

import (
	"log"
	"os"

	"github.com/go-mysql-org/go-mysql/client"
)

func Connection() *client.Conn {
	var server string = os.Getenv("MYSQL_SERVER")
	var database string = os.Getenv("MYSQL_DATABASE")
	var user string = os.Getenv("MYSQL_USER")
	var pass string = os.Getenv("MYSQL_PASSWORD")

	conn, err := client.Connect(server, user, pass, database)

	if err != nil {
		log.Fatal(err)
	}

	return conn
}
