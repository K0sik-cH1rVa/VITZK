package postre_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertOrUpdate(ctx context.Context, conn *pgx.Conn, id int, name string, age int, workhours int, post string) error {
	sqlQuery := `
	INSERT INTO users (id, name, age, workhours, post)
 	VALUES ($1, $2, $3, $4, $5)
 	ON CONFLICT (id) 
 	DO UPDATE SET name = EXCLUDED.name, post = EXCLUDED.post;`
	// При совпадении ID часы НЕ прибавляются, а просто обновляются имя/должность
	_, err := conn.Exec(ctx, sqlQuery, id, name, age, workhours, post)
	return err

}
