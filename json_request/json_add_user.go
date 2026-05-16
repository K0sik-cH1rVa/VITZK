package json_request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ychba/postre_sql"

	"github.com/jackc/pgx/v5"
)

// 1) Создаем структуру, в которую запишем данные
type Add_User struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	WorkHours int    `json:"work_hours"`
	Post      string `json:"post"`
}

func AddUser(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	// Проверяем, что это POST запрос
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	// в эту пустую переменную будет записываться текст который прислал клиент из POST запроса
	var newUser UserRequest

	// 2) парсим JSON прямо из тела запроса (r.Body) в нашу структуру
	// r.Body - тело входящего сетевого запроса
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Некорректный JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	// Отправляем данные в БД, где должна произойти запись нового пользователя
	ctx := r.Context()

	//Вставляем пользователя через InsertRow
	err := postre_sql.InsertRow(ctx, conn, newUser.Name, newUser.Age, newUser.WorkHours, newUser.Post)
	if err != nil {

		//Новый юзер! Вызываем ф-ию для его добавления
		err = postre_sql.Add_User(ctx, conn, newUser.Name, newUser.Age, newUser.WorkHours, newUser.Post)
		if err != nil {
			// Если это какая-то другая ошибка БД (упал сервер, нет таблицы и т.д.)
			http.Error(w, fmt.Sprintf("Ошибка сохранения в БД: %v", err), http.StatusInternalServerError)
			return
		}
		// Если ошибки не было — значит юзер создался впервые
		fmt.Printf("Пользователь %s успешно добавлен!\n", newUser.Name)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Пользователь успешно создан!"))
	}
	// ВОТ СЮДА В САМЫЙ НИЗ ВСТАВЛЯЙТЕ ЭТОТ БЛОК:
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Пользователь уже существует или произошла ошибка БД!"))
		return
	}

	// Вывод в консоль VS Code
	fmt.Printf("Успешно добавлен новый пользователь: %s\n", newUser.Name)

	// Отправка текста в Postman
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Добавлен новый юзер: %s", newUser.Name)))
} // Это самая последняя закрывающая скобка функции AddUser
