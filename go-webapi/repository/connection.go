package repository

import (
	"database/sql"
	"fmt"
)

func ConectaDb() *sql.DB {
	connStr := "user=golang password=golang dbname=godb host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("err")
		panic(err.Error())
	}
	return db
}
