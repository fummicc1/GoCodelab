package main

import (
	"database/sql"
	"log"
)

func main() {
	connection := "user=postgresuser dbname=postgresdatabase sslmode=verify-full"

	var db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
}
