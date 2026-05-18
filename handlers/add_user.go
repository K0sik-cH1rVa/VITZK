package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ychba/database"

	"github.com/jackc/pgx/v5/pgxpool"
)

// AddUserRequest - структура для данных от клиента
type AddUserRequest struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	WorkHours int    `json:"work_hours"`
	Post      string `json:"post"`
}

func AddUser(w http.ResponseWriter, r *http.Request, pool *pgxpool.Pool) {
	// Проверяем метод запроса
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST запросы", http.StatusMethodNotAllowed)
		return
	}

	// Читаем JSON из запроса
	var newUser AddUserRequest

	// Декодируем JSON
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Неправильный JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Проверяем возраст
	if newUser.Age < 0 || newUser.Age > 150 { //Пропущен оператор `||` (ИЛИ)
		http.Error(w, "Невалидный возраст", http.StatusBadRequest)
		return
	}

	// Проверяем имя
	if len(newUser.Name) == 0 || len(newUser.Name) > 100 { // Пропущен оператор `||` (ИЛИ)
		http.Error(w, "Имя пустое или слишком длинное", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	// Добавляем пользователя в БД
	err := database.InsertRow(ctx, pool, newUser.Name, newUser.Age, newUser.WorkHours, newUser.Post)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка БД: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Пользователь %s успешно добавлен!\n", newUser.Name)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Пользователь создан"))
}
