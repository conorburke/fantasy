package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var db *sql.DB

func init() {
	fmt.Println("connecting to database")
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("user=postgres password=%s host=database port=5432 dbname=postgres sslmode=disable", os.Getenv("FANTASY_DB_PASS")))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connection successful")
		// fmt.Println(db)
	}
}
