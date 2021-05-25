package data

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	_ "github.com/lib/pq"
)

var DBConn *sql.DB

const(
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "renol114"
	db = "A.E.R"
)

func OpenDBConnection() error{
	var err error
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, db)

	DBConn, err = sql.Open("postgrs", connectionString)

	return err
}
