package services

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct{
	DBConn *sql.DB
}


func (d *DB) OpenDBConnection(user, password, host, db, port string) error {
	var err error

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, db)

	d.DBConn, err = sql.Open("postgres", connStr)

	return err
}

func (d *DB) CloseDBConnection(){
	d.DBConn.Close()
}
