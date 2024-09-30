package db

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func ConnectToDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SaveRate(db *sql.DB, ask, bid float64, timestamp time.Time) error {
	_, err := db.Exec("INSERT INTO rates (ask, bid, timestamp) VALUES ($1, $2, $3)", ask, bid, timestamp)
	return err
}
