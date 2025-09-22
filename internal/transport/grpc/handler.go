package grpc

import (
	"context"
	"log"

	userpb "github.com/valdevay/project-protos/proto/user"
	"github.com/valdevay/users-service/internal/user"
)

// Handler implements the gRPC UserService
type Handler struct {
	svc *user.Service
	userpb.UnimplementedUserServiceServer
}

// NewHandler creates a new gRPC handler
func NewHandler(svc *user.Service) *Handler {
	return &Handler{svc: svc}
}

// CreateUser creates a new user
func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	log.Printf("CreateUser request: email=%s", req.Email)

	// Create user in service layer
	createdUser, err := h.svc.CreateUser(req.Email, "default_password")
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	// Convert to protobuf response
	response := &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:    uint32(createdUser.ID),
			Email: createdUser.Email,
		},
	}

	log.Printf("Created user: ID=%d, Email=%s", createdUser.ID, createdUser.Email)
	return response, nil
}

// GetUser retrieves a user by ID
func (h *Handler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	log.Printf("GetUser request: id=%d", req.Id)

	// Get user from service layer
	user, err := h.svc.GetUserByID(uint(req.Id))
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return nil, err
	}

	// Convert to protobuf response
	response := &userpb.GetUserResponse{
		User: &userpb.User{
			Id:    uint32(user.ID),
			Email: user.Email,
		},
	}

	log.Printf("Retrieved user: ID=%d, Email=%s", user.ID, user.Email)
	return response, nil
}

// UpdateUser updates an existing user
func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	log.Printf("UpdateUser request: id=%d, email=%s", req.User.Id, req.User.Email)

	// Update user in service layer
	updatedUser, err := h.svc.UpdateUser(uint(req.User.Id), req.User.Email, "")
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return nil, err
	}

	// Convert to protobuf response
	response := &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:    uint32(updatedUser.ID),
			Email: updatedUser.Email,
		},
	}

	log.Printf("Updated user: ID=%d, Email=%s", updatedUser.ID, updatedUser.Email)
	return response, nil
}

// DeleteUser deletes a user by ID
func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	log.Printf("DeleteUser request: id=%d", req.Id)

	// Delete user from service layer
	err := h.svc.DeleteUser(uint(req.Id))
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return nil, err
	}

	// Convert to protobuf response
	response := &userpb.DeleteUserResponse{
		Success: true,
	}

	log.Printf("Deleted user with id: %d", req.Id)
	return response, nil
}

// ListUsers retrieves all users
func (h *Handler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	log.Printf("ListUsers request: limit=%d, offset=%d", req.Limit, req.Offset)

	// Get all users from service layer
	users, err := h.svc.GetAllUsers()
	if err != nil {
		log.Printf("Error listing users: %v", err)
		return nil, err
	}

	// Convert to protobuf response
	var pbUsers []*userpb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &userpb.User{
			Id:    uint32(user.ID),
			Email: user.Email,
		})
	}

	response := &userpb.ListUsersResponse{
		Users: pbUsers,
		Total: uint32(len(pbUsers)),
	}

	log.Printf("Retrieved %d users", len(pbUsers))
	return response, nil
}
