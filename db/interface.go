package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"

	_ "github.com/go-sql-driver/mysql"
)

// DBConfig holds the configuration for the database connection
type DBConfig struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// Database is an interface for database operations
type Database interface {
	Connect() error
	Close() error
	Query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

// NewDatabase creates a new database connection based on the configuration
func NewDatabase(config *DBConfig) (Database, error) {
	var db Database
	switch config.Driver {
	case "postgres":
		db = &PostgresDB{config: config}
	case "mysql":
		db = &MySQLDB{config: config}
	case "mongodb":
		db = &MongoDB{config: config}
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", config.Driver)
	}
	if err := db.Connect(); err != nil {
		return nil, err
	}
	return db, nil
}
