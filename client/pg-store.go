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

func InitPgDB() {
	//? Create Tables
	err := createUserTable()
	if err != nil {
		panic(err)
	}
}

func createUserTable() error {
	query := `create table if not exists public.user (
		id serial primary key,
		first_name varchar(50) not null,
		middle_name varchar(50),
		last_name varchar(50) not null,
		email varchar(200) not null,
		phone_number varchar(255),
		password varchar not null,
		location varchar,
		age int,
		gender varchar(20),
		created_at timestamp,
		updated_at timestamp,
		deleted boolean not null default false
	);`

	_, err := dbClient.Exec(query)
	return err
}
