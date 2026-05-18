package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// 1) тут мы будем принимать запросы от БД
// 2) начинаем не с id, а с name тк в postgreSQL у нас id !!!SERIAL!!! PRIMARY KEY(этот пункт заполнит БД сама)
// 3) сначала пишем INSERT INTO потом название таблицы куда хотим вставить новую запись и в () перечисляем названия столбцов  и пишем key_word "VALUES" и в () передаем значение для каждого столбца!
// 4) пишем умных долларов потому что потому!
func InsertRow(
	ctx context.Context,
	pool *pgxpool.Pool,
	Name string,
	Age int,
	WorkHours int,
	Post string) error {
	sqlQuery := `
	INSERT INTO users (name, age, workhours, post)
	VALUES($1, $2, $3, $4);`
	// 1) отдельно пишем названия столбцов
	// 2) эта строка берет conn, текст sqlQuery и передает туда переменные Name - $1 etc.
	// 3) БД выполняет эту команду, создавая физическую строку на диске
	_, err := pool.Exec(ctx, sqlQuery, Name, Age, WorkHours, Post)
	return err
}
