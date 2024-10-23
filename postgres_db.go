package util

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	User string
	Pass string
	Host string
	DB   string
}

func (pg *PostgresDB) Open() *sql.DB {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", pg.User, pg.DB, pg.Pass, pg.Host)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}
