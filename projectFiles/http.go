package pgse

import (
	"encoding/json"
	"log/slog"
	sqlp "mainMod/projectFiles/sql"
	"net/http"
	"strconv"
)

func CreateProductH(w http.ResponseWriter, r *http.Request) { // Хендлер создания продукта
	if r.Method != http.MethodPost { // Проверка на метод подключения
		slog.Error("ERROR:", sqlp.WrongRequestMethod, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(sqlp.WrongRequestMethod.Error()))
		return
	}

	var input sqlp.ProductSQL // Создаем переменную для входных данных

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil { // Декодируем входящий JSON
		slog.Error("ERROR:", err, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if err := sqlp.ValidateALL(input.Name, input.Description, input.Price, input.Quantity); err != nil { // Делаем валидацию всех входящих значений
		slog.Error("ERROR:", err, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	outputR, err := sqlp.CreateProduct(sqlp.CTX, sqlp.Conn, input.Name, input.Description, input.Price, input.Quantity) // Создаем продукт
	if err != nil {
		slog.Error("ERROR:", err, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	output, err := json.Marshal(outputR) // Запаковываем созданный продукт в JSON-ответ
	if err != nil {
		slog.Error("ERROR:", err, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated) // Статус - создан
	w.Write(output)                   // Отправляем ответ
}

func ChangeProductH(w http.ResponseWriter, r *http.Request) { // Хендлер изменения продукта
	if r.Method != http.MethodPatch { // Проверка на метод подключения
		slog.Error("ERROR:", sqlp.WrongRequestMethod, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(sqlp.WrongRequestMethod.Error()))
		return
	}

	querry := r.URL.Query()                                                     // Обозначаем querry для более удобного создания переменных
	idR, key, value := querry.Get("id"), querry.Get("key"), querry.Get("value") // Получаем querry параметры из запроса
	id, err := strconv.Atoi(idR)                                                // Переводим ID в int для корректного запроса
	if err != nil {
		slog.Error("ERROR:", err, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	check, err := sqlp.CheckIdExists(id, sqlp.Conn, sqlp.CTX) // Проверяем на существование продукта
	if err != nil {
		slog.Error("ERROR:", err, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if check == false {
		slog.Error("ERROR:", err, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(sqlp.NotExists.Error()))
		return
	}
	if err := sqlp.ValidateKeyValue(key, value); err != nil { // Проверяем значения на замену
		slog.Error("ERROR:", err, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if err := sqlp.ChangeProduct(sqlp.CTX, sqlp.Conn, id, key, value); err != nil { // Изменяем продукт через функцию
		slog.Error("ERROR:", err, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Успех!"))
}

func DeleteProductH(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete { // Проверяем метод подключения
		slog.Error("ERROR:", sqlp.WrongRequestMethod, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(sqlp.WrongRequestMethod.Error()))
		return
	}

	idR := r.URL.Query().Get("id") // Получение querry параметра

	if idR == "" { // Проверка на то, пустой ли querry параметр
		slog.Error("ERROR:", sqlp.NotExists, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(sqlp.NotExists.Error()))
		return
	}

	id, err := strconv.Atoi(idR) // Переводим querry параметр в int
	if err != nil {
		slog.Error("ERROR:", sqlp.WrongRequestMethod, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	check, err := sqlp.CheckIdExists(id, sqlp.Conn, sqlp.CTX) // Проверяем на существование продукта
	if err != nil {
		slog.Error("ERROR:", err, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !check {
		slog.Error("ERROR:", sqlp.NotExists, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(sqlp.NotExists.Error()))
		return
	}

	if err := sqlp.DeleteProduct(sqlp.CTX, sqlp.Conn, id); err != nil { // Удаляем продукт через функцию
		slog.Error("ERROR:", err, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Успех!"))
}

func GetProductH(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { // Проверяем метод подключения
		slog.Error("ERROR:", sqlp.WrongRequestMethod, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(sqlp.WrongRequestMethod.Error()))
		return
	}

	page := r.URL.Query().Get("page")
	if page == "" { // Получаем страницу с querry запроса
		slog.Error("ERROR:", sqlp.WrongPage, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(sqlp.WrongPage.Error()))
		return
	}

	products, err := sqlp.GetProduct(sqlp.CTX, sqlp.Conn, page)
	if err != nil { // Создаем ответный слайс
		slog.Error("ERROR:", err, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	output, err := json.Marshal(products)
	if err != nil { // Запихиваем ответный слайс в JSON
		slog.Error("ERROR:", err, "METHOD:", r.Method, "PATH:", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(output)
}
