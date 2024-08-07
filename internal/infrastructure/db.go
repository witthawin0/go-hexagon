// internal/infrastructure/db.go
package infrastructure

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(dataSourceName string) *sql.DB {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
