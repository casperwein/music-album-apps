package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func OpenDB(dsn string, setLimits bool) (*sql.DB, error) {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if setLimits {
		fmt.Println("setting limits")
		db.SetMaxOpenConns(5)
		db.SetMaxIdleConns(5)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CloseDB(db *sql.DB) {

	err := db.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
