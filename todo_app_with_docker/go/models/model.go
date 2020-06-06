package models

import (
	"database/sql"
	"log"
)

// Model is a base type. All Models inherit Model.
type Model struct {
	ID int `json:id`
}

var db *sql.DB

func init() {
	connection := "user=postgresuser dbname=sampledatabase sslmode=verify-full"

	var err error
	db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
}
