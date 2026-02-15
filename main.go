package main

import (
	"log/slog"
	initf "mainMod/projectFiles/initFunc"
)

func main() {
	initf.CreateLogger()

	if err := initf.Getenv(); err != nil { // Импортируем переменные из .env
		slog.Error("ERROR:", err)
		return
	}
	if err := initf.ConnectAndCreate(); err != nil { // Создаем подключение к базе данных и таблицу в ней
		slog.Error("ERROR:", err)
		return
	}
	if err := initf.StartHTTP(); err != nil { // Создаем хендлеры и запускаем HTTP сервер
		slog.Error("ERROR:", err)
		return
	}
}
