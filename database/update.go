// нейрона тк не шарю за умные $, хотя идея была не столь отдалена
package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func UpdateRowHours(ctx context.Context, pool *pgxpool.Pool, name string, hours int) error {
	sqlQuery := `
	UPDATE users
	SET workhours = workhours + $1
	WHERE name = $2;`
	_, err := pool.Exec(ctx, sqlQuery, hours, name)
	if err != nil {
		return err
	}
	return nil
}
