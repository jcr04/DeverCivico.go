// infrastructure/dbconfig.go
package infrastructure

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	connStr := "user=postgres password=12345 dbname=civico	 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	return db, err
}
