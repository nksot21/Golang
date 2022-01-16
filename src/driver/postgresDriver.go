package driver

//https://github.com/lib/pq
//C:\Program Files\PostgreSQL\14\pgAdmin 4\bin
//http://127.0.0.1:54640/browser/
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	SQL *sql.DB
}

var Postgres = &PostgresDB{}

func Connect(host, port, user, password, dbname string) *PostgresDB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	Postgres.SQL = db
	return Postgres
}
