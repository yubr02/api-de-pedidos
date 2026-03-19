package db

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/glebarez/sqlite"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB(dsn string) (*sql.DB, error) {
	if strings.HasPrefix(dsn, "sqlite3:") || strings.HasPrefix(dsn, "file:") {
		dbPath := strings.TrimPrefix(dsn, "sqlite3:")
		if dbPath == "" {
			dbPath = "api_pedidos.db"
		}
		db, err := sql.Open("sqlite", dbPath)
		if err != nil {
			return nil, err
		}
		return db, nil
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db ping: %w", err)
	}
	return db, nil
}

func InitSchema(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		fullname TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		customer_name TEXT NOT NULL,
		items TEXT NOT NULL,
		total REAL NOT NULL,
		status TEXT NOT NULL DEFAULT 'pending',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
		`INSERT OR IGNORE INTO users (email, password, fullname) VALUES ('admin@api.com', 'password', 'Administrador');`,
	}

	for _, q := range queries {
		if _, err := db.Exec(q); err != nil {
			return err
		}
	}
	return nil
}
