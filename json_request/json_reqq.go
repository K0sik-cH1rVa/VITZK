// нейрона-мать
package json_request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ychba/users"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// 1. Создаем структуру, в которую запишем данные.
// Теги json:"..." обязательны, чтобы Go понимал, какое поле JSON соответствует полю структуры.
type UserRequest struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	WorkHours int    `json:"work_hours"`
	Post      string `json:"post"`
	Update    bool   `json:"update"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	// Проверяем, что это POST запрос
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	// в эту пустую переменную будет записываться текст который прислал клиент из POST запроса
	var newUser UserRequest

	// 2. Декодируем(парсим) JSON прямо из тела запроса (r.Body) в нашу структуру
	// r.Body - тело входящего сетевого запроса
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Некорректный JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Отправляем данные в БД, где должно? произойти сложение рабочих часов
	ctx := r.Context()

	// Вы вставляете пользователя через обычный InsertRow
	err = users.InsertRow(ctx, conn, newUser.Name, newUser.Age, newUser.WorkHours, newUser.Post)

	if err != nil {
		// Проверяем код ошибки PostgreSQL (23505 — это дубликат уникального ключа/имени)
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {

			// ЮЗЕР СУЩЕСТВУЕТ! Вызываем функцию обновления часов с вашего экрана
			err = users.UpdateRowHours(ctx, conn, newUser.Name, newUser.WorkHours)
			if err != nil {
				http.Error(w, "Ошибка обновления часов: "+err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Printf("Пользователь %s уже был в базе. Часы успешно прибавлены!\n", newUser.Name)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Часы обновлены!"))
			return

		} else {
			// Если это какая-то другая ошибка БД (упал сервер, нет таблицы и т.д.)
			http.Error(w, "Ошибка сохранения в БД: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Если ошибки не было — значит юзер создался впервые
	fmt.Printf("Успешно создан НОВЫЙ пользователь: %s\n", newUser.Name)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Пользователь успешно создан!"))
}

//err = users.InsertRow(ctx, conn, newUser.Name, newUser.Age, newUser.WorkHours, newUser.Post)
//if err != nil {
// 		http.Error(w, "ошибка сохранения в БД"+err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	fmt.Printf("Успешно добавлен пользователь: %s\n", newUser.Name)
// 	// Отправляем ответ клиенту
// 	w.WriteHeader(http.StatusCreated)
// 	w.Write([]byte("Пользователь успешно обработан!"))
// //}
