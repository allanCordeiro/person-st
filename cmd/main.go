package main

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/AllanCordeiro/person-st/infra/database"
	"github.com/AllanCordeiro/person-st/infra/webserver"
)

func main() {
	db, err := sql.Open("postgres", "postgres://rinha:rinha123@db/rinhadb?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	personDB := database.NewPersonDB(db)

	webserver.Serve(personDB)
}
