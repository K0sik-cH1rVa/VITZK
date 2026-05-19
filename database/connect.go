package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// создаем подключение к БД и его валидность:
func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	dbURL := os.Getenv("DbURL")
	if dbURL == "" {
		fmt.Println("Пустой пароль")
	}
	return pgxpool.New(ctx, dbURL)
}
