package database

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func GetDBConnection(cfg *mysql.Config) *sql.DB {
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return db
}
