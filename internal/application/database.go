package application

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)

	// perf:
	db.SetMaxOpenConns(25) // максимум одновременных соединений
	db.SetMaxIdleConns(25) // сколько держать открытыми в idle
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
