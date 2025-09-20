package grpc

// gRPC handlers will be defined here
import (
	"context"
	userpb "github.com/valdevay/project-protos/proto/users"
	"github.com/valdevay/users-service/internal/user"
)

type Handler struct {
	svc *user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc *user.Service) *Handler {
	return &Handler{svc: svc}
}

// CreateUser создает нового пользователя
func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	// Преобразуем req → user.User
	userModel := &user.User{
		Email:    req.Email,
		Password: "default_password", // Временное решение, так как Password не доступен в protobuf
	}
	
	// Вызываем svc.CreateUser
	createdUser, err := h.svc.CreateUser(ctx, *userModel)
	if err != nil {
		return nil, err
	}
	
	// Преобразуем user.User → userpb.User
	userProto := &userpb.User{
		Id:    createdUser.ID, // Используем реальный ID пользователя
		Email: createdUser.Email,
	}
	
	return &userpb.CreateUserResponse{User: userProto}, nil
}

// GetUser получает пользователя по ID  
func (h *Handler) GetUser(ctx context.Context, req *userpb.User) (*userpb.User, error) {
	// Вызываем svc.GetUserByID
	userModel, err := h.svc.GetUserByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	
	// Преобразуем user.User → userpb.User
	userProto := &userpb.User{
		Id:    userModel.ID, // Используем реальный ID пользователя
		Email: userModel.Email,
	}
	
	return userProto, nil
}

// ListUsers получает список всех пользователей
func (h *Handler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	// Вызываем svc.GetAllUsers
	users, err := h.svc.GetAllUsers(ctx, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}
	
	// Преобразуем срез user.User → []*userpb.User
	usersProto := make([]*userpb.User, len(users))
	for i, userModel := range users {
		usersProto[i] = &userpb.User{
			Id:    userModel.ID, // Используем реальный ID пользователя
			Email: userModel.Email,
		}
	}
	
	return &userpb.ListUsersResponse{Users: usersProto}, nil
}

// UpdateUser обновляет пользователя
func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	// Преобразуем req.User → user.User
	userModel := &user.User{
		ID:    req.User.Id,
		Email: req.User.Email,
	}
	
	// Вызываем svc.UpdateUserByID
	updatedUser, err := h.svc.UpdateUserByID(ctx, req.User.Id, *userModel)
	if err != nil {
		return nil, err
	}
	
	// Преобразуем user.User → userpb.User
	userProto := &userpb.User{
		Id:    updatedUser.ID, // Используем реальный ID пользователя
		Email: updatedUser.Email,
	}
	
	return &userpb.UpdateUserResponse{User: userProto}, nil
}

// DeleteUser удаляет пользователя по ID
func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	// Вызываем svc.DeleteUserByID
	err := h.svc.DeleteUserByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	
	return &userpb.DeleteUserResponse{Id: req.Id}, nil
}