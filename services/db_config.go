package services

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DBConn *sql.DB


func OpenDBConnection(user, password, host, port, db string) error {
	var err error
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, db)


	DBConn, err = sql.Open("postgres", connStr)

	return err
}
