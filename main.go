package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/kaungmyathan22/golang-sqlc/database"
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
	err = db.Ping()
	if err != nil {
		log.Println("Something went wrong while pinging to database!!")
		panic(err)
	}
	ctx := context.Background()
	queries := database.New(db)
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		log.Fatal("error occur while fetching authors")
		panic(err)
	}
	for _, element := range authors {
		fmt.Println(element.ID, element.Name, element.Bio.String)
	}
	// queries.CreateAuthor(ctx, database.CreateAuthorParams{Name: "Kaung Myat Han", Bio: sql.NullString{String: "I am a developer.", Valid: true}})
	// author, err := queries.GetAuthor(ctx, 1)
	// if err != nil {
	// 	log.Printf("error occur while fetching author with id %d\n", 1)
	// 	panic(err)
	// }
	// err = queries.UpdateAuthor(ctx, database.UpdateAuthorParams{ID: author.ID, Bio: sql.NullString{String: "Updated Bio", Valid: true}, Name: "Eric Han"})
	// if err != nil {
	// 	log.Printf("error occur while updating author with id %d\n", 1)
	// 	panic(err)
	// }
	err = queries.DeleteAuthor(ctx, 1)
	if err != nil {
		log.Printf("error occur while deleting author with id %d\n", 1)
		panic(err)
	} else {
		log.Println("successfully deleted author")
	}
	log.Println("successfully connected to database!!!")
}
