package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type MySQLDB struct {
	config *DBConfig
	db     *sqlx.DB
}

func (m *MySQLDB) Connect() error {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		m.config.User, m.config.Password, m.config.Host, m.config.Port, m.config.DBName)
	db, err := sqlx.Connect("mysql", connStr)
	if err != nil {
		return err
	}
	m.db = db
	return m.db.Ping()
}

func (m *MySQLDB) Close() error {
	return m.db.Close()
}

func (m *MySQLDB) Query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	return m.db.QueryxContext(ctx, query, args...)
}

func (m *MySQLDB) QueryRow(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return m.db.QueryRowxContext(ctx, query, args...)
}

func (m *MySQLDB) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return m.db.ExecContext(ctx, query, args...)
}
