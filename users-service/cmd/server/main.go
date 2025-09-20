package main

import (
	"log"

	"github.com/valdevay/users-service/internal/database"
	"github.com/valdevay/users-service/internal/user"
	transportgrpc "github.com/valdevay/users-service/internal/transport/grpc"
)

func main() {
	database.InitDB()
	repo := user.NewRepository(database.DB)
	svc := user.NewService(repo)

	if err := transportgrpc.RunGRPC(svc); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}
}
