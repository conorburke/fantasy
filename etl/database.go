package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	fmt.Println("connecting to database")
	var err error
	db, err = sql.Open("postgres", "user=postgres password=postgres host=database port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connection successful")
		fmt.Println(db)
	}
}
