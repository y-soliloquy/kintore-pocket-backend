package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DBClient struct {
	DB *sql.DB
}

func NewDBClient(DatabaseURL string) (*DBClient, error) {
	db, err := sql.Open("mysql", DatabaseURL)
	if err != nil {
		return nil, err
	}

	// プール設定（環境共通）
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return &DBClient{
		DB: db,
	}, nil
}

func (c *DBClient) Close() error {
	return c.DB.Close()
}
