// нейрочип мне в анус!
package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetAllUsers(ctx context.Context, pool *pgxpool.Pool) ([]user, error) {
	sqlQuery := `
	SELECT id, name, age, workhours, post FROM users;`
	// метод Query отправляет текст запроса в Postgres, БД отдает его в rows
	rows, err := pool.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// allUsers - слайс-коробка и чтение через rows.Next()
	var allUsers []user

	// Запускает цикл который идет по строкам в БД...
	for rows.Next() {
		var u user // Временная пустая структура для одной строки

		// Сканируем данные из БД в поля структуры через знак &
		err := rows.Scan(&u.Id, &u.Name, &u.Age, &u.WorkHours, &u.Post)
		if err != nil {
			return nil, err // Если сканирование строки сломалось - return
		}

		// Добавляем заполненного юзера в общий слайс
		allUsers = append(allUsers, u)
	}

	// Финальная проверка не упало ли соединение во время работы цикла
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Возвращаем готовый массив пользователей
	return allUsers, nil
}
