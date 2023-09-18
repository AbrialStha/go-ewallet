package client

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	dbClient *sql.DB
)

func PgConnect() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "app-user"
		password = "mysecret123"
		dbname   = "go-ewallet-db"
	)

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	dbClient = db
}

func GetPgDB() *sql.DB {
	return dbClient
}
