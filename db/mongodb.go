package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	config *DBConfig
	client *mongo.Client
}

func (m *MongoDB) Connect() error {
	connStr := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
		m.config.User, m.config.Password, m.config.Host, m.config.Port, m.config.DBName)
	clientOptions := options.Client().ApplyURI(connStr)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	m.client = client
	return m.client.Ping(context.Background(), nil)
}

func (m *MongoDB) Close() error {
	return m.client.Disconnect(context.Background())
}

func (m *MongoDB) Query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *MongoDB) QueryRow(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return nil
}

func (m *MongoDB) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return nil, fmt.Errorf("not implemented")
}
