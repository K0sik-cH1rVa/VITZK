package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"ychba/database"
	"ychba/handlers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load() // Загрузит .env
	if err != nil {
		log.Println("Файл .env не найден, используем системные переменные")
	}

	ctx := context.Background()

	// 1. Подключаемся к БД и создаем пул соединений
	pool, err := database.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	// 2. Создаем таблицу с юзерами
	err = database.CreateTable(ctx, pool)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Успешно подключились к БД :)")

	// 3. Записываем ВСЕХ наших имеющихся юзеров
	// err = users.InsertAllUsers(ctx, pool, users.Users)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 4. Настраиваем http-эндпоинты
	http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		handlers.AddUser(w, r, pool)
	})

	http.HandleFunc("/api/users/hours", func(w http.ResponseWriter, r *http.Request) {
		handlers.AddWorkHours(w, r, pool)
	})

	// 5. Запуск сервера
	fmt.Println("Запускаю http сервер на порту :8080...")
	svc := &http.Server{Addr: ":8080"}

	if err := svc.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Сервер упал: ", err)
	}
}
