package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/IsaacDSC/desafio-padawan-go/src/infra/environments"
	_ "github.com/go-sql-driver/mysql"
)

func GetConnectionMysql() *sql.DB {
	env := environments.GetEnv()
	DATABASE_URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", env.MYSQL_USER, env.MYSQL_PASS, env.MYSQL_HOST, env.MYSQL_PORT, env.MYSQL_DATABASE)
	db, err := sql.Open("mysql", DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
