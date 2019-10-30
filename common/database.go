package common

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	fmt.Println("connecting to database")
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=postgres host=database port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connection successful")
		fmt.Println(DB)
	}
}
