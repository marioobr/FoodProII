package driver

import "sync"

const (
	host = "localhost"
	port = "5432"
	user = "postgres"
	pwd  = "Ab321"
	db   = "foodPro"
)

var connect *PostgresDB
var sincro sync.Once

func Instance() *PostgresDB {
	sincro.Do(func() {
		connect = Manager()
	})
	return connect
}

func Manager() *PostgresDB {
	database := Connect(host, port, user, pwd, db)
	return database
}
