package main

import (
	"log"

	"github.com/valdevay/users-service/internal/database"
	transportgrpc "github.com/valdevay/users-service/internal/transport/grpc"
	"github.com/valdevay/users-service/internal/user"
)

func main() {
	// Инициализация базы данных
	database.InitDB()

	// Инициализация репозитория и сервиса
	repo := user.NewRepository(database.DB)
	svc := user.NewService(repo)

	// Запуск gRPC сервера
	if err := transportgrpc.RunGRPC(svc); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}
}
