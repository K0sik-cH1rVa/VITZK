package postre_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// создаем подключение к БД и его валидность:
func Connect(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "postgres://postgres:22853@localhost:5432/postgres")
}
