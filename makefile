start: # Для запуска ProductManager без компиляции и базы данных (нужна внешняя)
	go run main.go 

deploy-up: # Для запуска ProductManager
	docker compose up -d pmanager-service

deploy-down: # Для остановки ProductManager
	docker-compose down
