package initf

import (
	"log/slog"
	"os"
)

func CreateLogger() {
	logFile, err := os.Create("/var/log/pmanager.log") // Создаем лог файл
	if err != nil {
		panic("Не удалось создать файл лога: " + err.Error())
	}
	defer logFile.Close()

	logger := slog.New(slog.NewTextHandler(logFile, nil)) // Создаём новый логгер с выводом в файл
	slog.SetDefault(logger)                               // Делаем его по умолчанию

	slog.Info("Сервер запущен")
}
