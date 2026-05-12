package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() {
	//создаем экземпляр контекста и положим в переменную
	ctxx := context.Background()

	conn, err := pgx.Connect(ctxx, "postgres://postgres:22853@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	//это тестовый запрос в БД, который проверяет его валидность
	if err := conn.Ping(ctxx); err != nil {
		panic(err)
	}
	//Если мы на этом моменте, то мы упешно создали валидное подключение и смогли проверить его пинг запросом conn.Ping(ctx), в таком случае мы выведем :
	fmt.Println("Подключение к базе данных прошло успешно!")
}
