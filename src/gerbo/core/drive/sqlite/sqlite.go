// Gerbo - Rodent and data extractor
// https://github.com/luk4z7/gerbo for the canonical source repository
// Copyright Lucas Alves 2017

// core/drive
package sqlite

import (
	"database/sql"
	_ "rsc.io/sqlite"
	"log"
)

const (
	drive    = "sqlite3"
	database = "/data/twitter-movie-ratings.db"
)

func GetDB() *sql.DB {
	db, err := sql.Open(drive, database)
	if err != nil {
		log.Fatalln("could not open database:", err)
	}
	return db
}