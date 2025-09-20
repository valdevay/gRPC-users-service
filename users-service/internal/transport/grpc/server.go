package grpc

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	userpb "github.com/valdevay/project-protos/proto/users"
	"github.com/valdevay/users-service/internal/user"
)

// RunGRPC настраивает и запускает gRPC-сервер
func RunGRPC(svc *user.Service) error {
	// 1. Создаем listener на порту 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to listen on port 50051: %w", err)
	}

	// 2. Настраиваем параметры gRPC сервера
	keepaliveParams := keepalive.ServerParameters{
		Time:    10 * time.Second, // Отправляем ping каждые 10 секунд
		Timeout: 5 * time.Second,  // Ждем pong 5 секунд
	}

	// 3. Создаем gRPC сервер с настройками
	grpcSrv := grpc.NewServer(
		grpc.KeepaliveParams(keepaliveParams),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             5 * time.Second, // Минимальный интервал между ping'ами
			PermitWithoutStream: true,            // Разрешаем ping'и без активных стримов
		}),
	)

	// 4. Регистрируем UserService
	userpb.RegisterUserServiceServer(grpcSrv, NewHandler(svc))

	// 5. Логируем запуск сервера
	log.Printf("Starting gRPC server on port 50051...")
	
	// 6. Запускаем сервер
	if err := grpcSrv.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

// TODO: Добавить graceful shutdown для корректного завершения работы сервера
// TODO: Добавить health check endpoint для мониторинга
// TODO: Добавить метрики и логирование запросов
// TODO: Добавить TLS конфигурацию для production
// TODO: Добавить middleware для аутентификации и авторизации