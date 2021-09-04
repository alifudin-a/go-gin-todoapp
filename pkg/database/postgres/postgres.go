package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var PG *sqlx.DB

func OpenPG() {

	var (
		host   = os.Getenv("PG_DB_HOST")
		port   = os.Getenv("PG_DB_PORT")
		user   = os.Getenv("PG_DB_USER")
		pass   = os.Getenv("PG_DB_PASS")
		dbname = os.Getenv("PG_DB_NAME")
	)

	pgInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)

	db, err := sqlx.Open("postgres", pgInfo)
	if err != nil {
		log.Println("Error while connecting to database: ", err)
	} else {
		log.Printf("Successfully connected to database '%s'!", dbname)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Minute * 5)

	PG = db
}
