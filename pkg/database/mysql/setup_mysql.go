package database

import (
    "database/sql"
    "log"
    "github.com/go-sql-driver/mysql"
)

func SetupMySQLDatabase(cfg mysql.Config) (*sql.DB, error) {
	dsn := cfg.FormatDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("ERROR OPENING DATABASE: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("ERROR PINGING DATABASE: %v", err)
	}

	log.Println("DATABASE CONNECTION SUCCESSFUL")
	return db, nil
}

