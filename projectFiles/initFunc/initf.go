package initf

import (
	"fmt"
	pgse "mainMod/projectFiles"
	sqlp "mainMod/projectFiles/sql"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Getenv() error { // Инициализируем переменные из .env
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}

func ConnectAndCreate() error { // Подключиться к базе данных и создать таблицу
	var err error

	var dbUser string = os.Getenv("DB_USER") // Получаем переменные из .env
	var dbPass string = os.Getenv("DB_PASSWORD")
	var dbName string = os.Getenv("DB_NAME")
	var dbHost string = os.Getenv("DB_HOST")
	var dbPort string = os.Getenv("DB_PORT")
	var connection string = "postgres://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName
	sqlp.Conn, err = sqlp.CreateConnection(sqlp.CTX, connection) // Подключаемся к базе данных

	if err != nil { // Проверяем на наличие ошибок
		return err
	}
	if err := sqlp.CreateTable(sqlp.CTX, sqlp.Conn); err != nil { // Создаем таблицу в базе данных
		return err
	}
	return nil
}

func StartHTTP() error { // Создаем каталоги и запускаем HTTP сервер
	http.HandleFunc("/create", pgse.CreateProductH)
	http.HandleFunc("/change", pgse.ChangeProductH)
	http.HandleFunc("/delete", pgse.DeleteProductH)
	http.HandleFunc("/", pgse.GetProductH)

	fmt.Println("HTTP Сервер запущен! | HTTP Server started!") // Индикация старта сервера в консоль
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return err
	}
	return nil
}
