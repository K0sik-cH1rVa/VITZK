package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// 1) в этой ф-ии в ее аргументы также добавляем наш контекст, тк ф-ия Exec первым аргументом принимает контекст
func CreateTable(ctx context.Context, pool *pgxpool.Pool) error {
	sqlQuery := `
CREATE TABLE IF NOT EXISTS users(
id SERIAL PRIMARY KEY,
name VARCHAR(100) NOT NULL UNIQUE,
age INT,
workhours INT NOT NULL,
post VARCHAR(50) NOT NULL
);`
	//неправильно - тип на будущее
	//pool, err := pgxpool.New(ctx, sqlQuery) //- тут я пытаюсь подключиться к SQL запросу, а надо его выполнить
	//return err
	_, err := pool.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}
	return nil
}
