package postre_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// 1) в этой ф-ии в ее аргументы также добавляем наш контекст, тк ф-ия Exec первым аргументом принимает контекст
func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users(
id SERIAL PRIMARY KEY,
name VARCHAR(100) NOT NULL UNIQUE,
age INT,
workHours INT NOT NULL,
post VARCHAR(50) NOT NULL
);`
	_, err := conn.Exec(ctx, sqlQuery)
	return err
}
