package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"ychba/json_request"
	"ychba/postre_sql"
	"ychba/users"
)

// глобальная переменная сервера
var svc *http.Server

func Handler(w http.ResponseWriter, r *http.Request) {
	//нейрона//
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(users.Users)
} //end//
// ДАТТЕБАЁ
func main() {
	// создаем контекст Background:
	ctx := context.Background()
	//Connect принимает контекст и возвращает созданное подключени к БД/ошибку
	conn, err := postre_sql.Connect(ctx)
	if err != nil {
		panic(err)
	}
	//создаем таблицу с юзерами
	err = postre_sql.CreateTable(ctx, conn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Успешно подключились к БД :)")
	//записываем ВСЕХ наших имеющихся юзеров - нейрона(не знал как закинуть всех+)
	err = users.InsertAllUsers(ctx, conn, users.Users)
	if err != nil {
		panic(err)
	}
	//1) - выводит на экран по ссылке в браузере 3 юзера а остальных новых в бд стирает :(
	http.HandleFunc("/anal", Handler)

	//2/3) норм ребята, еще и в консоли приятность и вкусность делают(помимо закидывания типов в БД)
	http.HandleFunc("/api/addUser/addWorkHours", func(w http.ResponseWriter, r *http.Request) {
		json_request.AddWorkHours(w, r, conn)
	})
	http.HandleFunc("/api/addUser", func(w http.ResponseWriter, r *http.Request) {
		json_request.AddUser(w, r, conn)
	})

	// 2) И только после этого финальный запуск сервера:
	fmt.Println("Запускаю http сервер на порту :8080...")
	//настраиваем http-сервер//
	//инициализируем глобальную переменную чтобы != nil//
	svc = &http.Server{Addr: ":8080"}

	//запускаем горутину - нейрона
	go postre_sql.CheckID(ctx, conn)

	//Запускаем сервер через эту переменную//
	if err := svc.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Println("Произошла ошибка :(", err.Error())
		return
	}
}
