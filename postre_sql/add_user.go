package postre_sql
import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Add_User(ctx context.Context, conn *pgx.Conn, name string, age int, workhours int, post string) error {
	sqlQuery := `
	INSERT INTO users (name, age, workhours, post)
	VALUES ($1, $2, $3, $4);`
	_, err := conn.Exec(ctx, sqlQuery, name, age, workhours, post)
	return err
}
