package databaseutils

import (
	"config"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func init() {
	dbtmp, err := sql.Open("postgres", config.GetDbConnStr())
	if err != nil {
		log.Printf("Can't connect to database %v\n", err)
	}
	db = dbtmp
	log.Println("Connected sucsessfully to database (but don't hope too much, it still can fall :) )")
}

func GetDB() *sql.DB {
	return db
}
