// нейрона тк не шарю за умные $, хотя идея была не столь отдалена
package postre_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRowHours(ctx context.Context, conn *pgx.Conn, name string, hours int) error {
	sqlQuery := `
	UPDATE users
	SET workhours = workhours + $1
	WHERE name = $2;`
	_, err := conn.Exec(ctx, sqlQuery, hours, name)
	return err
}
