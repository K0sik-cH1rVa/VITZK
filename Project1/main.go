package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	Name      string
	Age       int
	ID        int
	WorkHours int
	Post      string
}

var userr1 = user{
	Name:      "woeg",
	Age:       13,
	ID:        1,
	WorkHours: 5,
	Post:      "работяга",
}

var userr2 = user{
	Name:      "bebra",
	Age:       18,
	ID:        2,
	WorkHours: 1000,
	Post:      "Admin",
}

var userr3 = user{
	Name:      "антон",
	Age:       16,
	ID:        3,
	WorkHours: 45,
	Post:      "нищета",
}

var userss = []user{
	userr1,
	userr2,
	userr3,
}

// глобальная переменная сервера
var svc *http.Server

func Handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(userss)
	go func() {
		if svc != nil {
			_ = svc.Shutdown(context.Background())
		}
	}()
}

func main() {

	http.HandleFunc("/anal", Handler)
	fmt.Println("Запускаю http сервер!")
	//настраиваем http-сервер//
	//инициализируем глобальную переменную чтобы != nil//
	svc = &http.Server{Addr: ":8080"}
	//Запускаем сервер через эту переменную
	if err := svc.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Println("Произошла ошибка", err.Error())
		return
	}
	fmt.Println("сервер успешно остановлен!")

	// создаем контекст Background:
	ctx := context.Background()
	//CheckConnection принимает контекст и возвращает созданное подключеник к БД/ошибку
	if _, err := Connect(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Успешно подключились к БД!")
}
