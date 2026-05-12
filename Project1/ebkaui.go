package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Добавляем методы в интерфейс
type Usser interface {
	GetName() string
	GetAge() int
	GetID() int
	GetWorkHours() int
	GetPost() string
}
type User struct {
	Name      string
	Age       int
	ID        int
	WorkHours int
	Post      string
}

// Пишем всех юзеров через var тк вне ф-ий нельзя кратко обьявлять переменные!!!, но есть способ чище т это я закоментирую.
var user1 = User{
	Name:      "bebra",
	Age:       18,
	ID:        1947240,
	WorkHours: 1000,
	Post:      "Admin",
}
var user2 = User{
	Name:      "bebra2",
	Age:       17,
	ID:        1894730,
	WorkHours: 46,
	Post:      "Работяга",
}
var user3 = User{
	Name:      "bebra3",
	Age:       16,
	ID:        2943920,
	WorkHours: 24,
	Post:      "Нищета",
}

// Заносим всех в один слайс (массив) интерфейсов
var users = []Usser{
	user1,
	user2,
}

// 2 пишем именованную ф-ию
func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		fmt.Println("Произошла ошибка!", err.Error())
	}
}

// Реализуем методы для нашей структуры
func (u User) GetName() string {
	return u.Name
}
func (u User) GetAge() int {
	return u.Age
}
func (u User) GetID() int {
	return u.ID
}
func (u User) GetWorkHours() int {
	return u.WorkHours
}
func (u User) GetPost() string {
	return u.Post
}

func main() {
	//можно удобно перебирать их в цикле
	for _, u := range users {
		fmt.Println("Имя пользователя:", u.GetName())
		fmt.Println("Возраст пользователя:", u.GetAge())
		fmt.Println("ID пользователя:", u.GetID())
	}

	http.HandleFunc("/analn", handler)
	//Пишем nil тк мы все настроили ранее и эта бурмалда нужна для "более крутой настройки бурмалды"
	// И тк эта ф-ия возвращает ошибку, мы будем ее выводить:
	fmt.Println("Запускаю http сервер!")
	err := http.ListenAndServe(":5432", nil)
	if err != nil {
		fmt.Println("Произошла ошибка", err.Error())
	} else {
		fmt.Println("Я корректно получил http запрос!")
	}
	fmt.Println("Программа завершила свое выполнение!")
	CheckConnection()
}
