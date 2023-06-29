package repository

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	c := mysql.Config{
		DBName: os.Getenv("MYSQL_DATABASE"),
		User: os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Addr: "db:3306",
		Net: "tcp",
	}
	Db, err = sql.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Print(err)
	}
}