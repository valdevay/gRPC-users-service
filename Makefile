.PHONY: build run clean deps test

# Установка зависимостей
deps:
	go mod download
	go mod tidy

# Сборка приложения
build:
	go build -o bin/users-service ./cmd/server

# Запуск приложения
run: build
	./bin/users-service

# Тестирование (запуск с тестами)
test: build
	./bin/users-service

# Очистка
clean:
	rm -rf bin/

# Запуск PostgreSQL в Docker
postgres:
	docker run --name users-postgres -e POSTGRES_DB=users -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres:13

# Остановка PostgreSQL
stop-postgres:
	docker stop users-postgres
	docker rm users-postgres
