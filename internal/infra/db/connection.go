package db

import (
	"database/sql"
)

func NewSQLiteConnection(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Initialize schema
	if err = initSchema(db); err != nil {
		return nil, err
	}

	return db, nil
}

func initSchema(db *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS cafes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		rating REAL,
		reviews INTEGER,
		price_range TEXT,
		type TEXT,
		address TEXT,
		review_text TEXT
	);`

	_, err := db.Exec(createTableSQL)
	return err
}
