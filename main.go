package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	const postgresURL = "postgres://admin:admin@localhost:5432/sqlc?sslmode=disable"
	db, err := sql.Open("postgres", postgresURL)
	defer db.Close()
	if err != nil {
		log.Println("Something went wrong while connecting to database!!")
		panic(err)
	}
}
