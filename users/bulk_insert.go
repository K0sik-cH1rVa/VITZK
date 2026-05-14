// нейрона мать + комменты
package users

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// InsertAllUsers принимает весь слайс пользователей и отправляет в БД одним запросом
func InsertAllUsers(ctx context.Context, conn *pgx.Conn, allUsers []user) error {
	// Создаем таблицу (срез срезов) для массовой вставки
	var rows [][]any

	for _, u := range allUsers {
		// Порядок полей должен строго совпадать с колонками в базе данных
		rows = append(rows, []any{u.Name, u.Age, u.WorkHours, u.Post})
	}

	// Команда CopyFrom отправляет все данные за один сетевой пакет
	_, err := conn.CopyFrom(
		ctx,
		pgx.Identifier{"users"}, // название таблицы
		[]string{"name", "age", "workhours", "post"}, // Колонки в таблице
		pgx.CopyFromRows(rows),                       // Данные для вставки
	)

	return err
}
