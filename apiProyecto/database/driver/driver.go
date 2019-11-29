package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	SQL *sql.DB
}

var Postgres = &PostgresDB{}

func Connect(host, port, user, password, dbname string) *PostgresDB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	Postgres.SQL = db
	return Postgres
}

func(p *PostgresDB) Close() error {
	return p.Close()
}