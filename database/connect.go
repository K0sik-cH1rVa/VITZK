package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// создаем подключение к БД и его валидность:
func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	dbURL := os.Getenv("DbURL")
	if dbURL == "" {
		dbURL = "postgres://postgres:22853@localhost:5432/postgres" // Значение по умолчанию
	}
	return pgxpool.New(ctx, dbURL)
}
