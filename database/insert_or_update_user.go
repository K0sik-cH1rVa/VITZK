package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertOrUpdate(ctx context.Context, pool *pgxpool.Pool, id int, name string, age int, workhours int, post string) error {
	sqlQuery := `
	INSERT INTO users (id, name, age, workhours, post)
 	VALUES ($1, $2, $3, $4, $5)
 	ON CONFLICT (id) 
 	DO UPDATE SET name = EXCLUDED.name, post = EXCLUDED.post;`
	// При совпадении ID часы НЕ прибавляются, а просто обновляются имя/должность
	pool, err := pgxpool.New(ctx, sqlQuery)
	return err

}
