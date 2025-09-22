package grpc

import (
	"fmt"
	"log"
	"net"

	userpb "github.com/valdevay/project-protos/proto/user"
	"github.com/valdevay/users-service/internal/user"
	"google.golang.org/grpc"
)

// RunGRPC starts the gRPC server
func RunGRPC(svc *user.Service) error {
	// Listen on port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register UserService
	userpb.RegisterUserServiceServer(grpcServer, NewHandler(svc))

	log.Println("gRPC server starting on :50051")
	log.Println("Available methods:")
	log.Println("  - CreateUser")
	log.Println("  - GetUser")
	log.Println("  - UpdateUser")
	log.Println("  - DeleteUser")
	log.Println("  - ListUsers")

	// Start serving
	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("gRPC server failed: %v", err)
	}

	return nil
}
