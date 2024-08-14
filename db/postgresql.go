package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PostgresDB struct {
	config *DBConfig
	db     *sqlx.DB
}

func (p *PostgresDB) Connect() error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		p.config.Host, p.config.Port, p.config.User, p.config.Password, p.config.DBName, p.config.SSLMode)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return err
	}
	p.db = db
	return p.db.Ping()
}

func (p *PostgresDB) Close() error {
	return p.db.Close()
}

func (p *PostgresDB) Query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	return p.db.QueryxContext(ctx, query, args...)
}

func (p *PostgresDB) QueryRow(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return p.db.QueryRowxContext(ctx, query, args...)
}

func (p *PostgresDB) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return p.db.ExecContext(ctx, query, args...)
}
