package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	// postgres driver
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

func InitDB() {
	// Use root:dbpass@tcp(172.17.0.2)/hackernews, if you're using Windows.
	db, err := sqlx.Open("postgres", "dbname=gunpla sslmode=disable")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	Db = db
}
