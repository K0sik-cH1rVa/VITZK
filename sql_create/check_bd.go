package sql_create

import (
	"context"
	"fmt"
	"time"
	"ychba/users"

	"github.com/jackc/pgx/v5"
)

// фоновая ф-ия мониторинга базы данных
func CheckBd(ctx context.Context, conn *pgx.Conn) {
	//создаем таймер на 5 минут
	tiki := time.Tick(25 * time.Second)
	//бесконечный цикл проверки:
	for range tiki {
		fmt.Println("проверка пошла нах")
		//вставляем всех типси
		err := users.InsertAllUsers(ctx, conn, users.Users)
		if err != nil {
			fmt.Println("Ошибка", err.Error())
		} else {
			fmt.Println("Проверка завершена, новые пользователи добавлены, существующие пропущены.")
		}
	}

}
