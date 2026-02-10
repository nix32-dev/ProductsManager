package main

import (
	"fmt"
	pgse "mainMod/projectFiles"
	sqlp "mainMod/projectFiles/sql"
	"net/http"
)

func main() {
	var err error
	sqlp.Conn, err = sqlp.CreateConnection(sqlp.CTX, "postgres", "pass", "localhost:5432", "postgres") // Подключаемся к базе данных | Введите ваш: user, password, ip, db name
	if err != nil {
		fmt.Println("Ошибка: | Error: ", err)
		return
	}
	if err := sqlp.CreateTable(sqlp.CTX, sqlp.Conn); err != nil { // Создаем таблицу в базе данных
		fmt.Println("Ошибка: | Error: ", err)
	}

	http.HandleFunc("/create", pgse.CreateProductH)
	http.HandleFunc("/change", pgse.ChangeProductH)
	http.HandleFunc("/delete", pgse.DeleteProductH)
	http.HandleFunc("/", pgse.GetProductH)

	fmt.Println("Старт!") // Индикация старта сервера в консоль
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка: ", err) // Проверка на ошибки, их вывод
	}
}
