package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnectionMysql() *sql.DB {
	db, err := sql.Open("mysql", "dev:secret@/desafiopadawango")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
