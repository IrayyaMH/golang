package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func Connect() {
	// Please define your user name and password for my sql.
	d, err := sql.Open("mysql", "username:Irayya@1993@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *sql.DB {
	return db
}
