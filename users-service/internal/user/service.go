package user

import (
	"context"
	"crypto/rand"
	"encoding/hex"
)

type Service struct {
	repo UserRepository
}

func NewService(repo UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllUsers(ctx context.Context, limit, offset int) ([]User, error) {
	return s.repo.GetAllUsers(ctx, limit, offset)
}

func (s *Service) CreateUser(ctx context.Context, user User) (User, error) {
	// Генерируем уникальный ID
	user.ID = generateID()
	return s.repo.CreateUser(ctx, user)
}

func (s *Service) UpdateUserByID(ctx context.Context, id string, user User) (User, error) {
	user.ID = id // Убеждаемся, что ID не изменится
	return s.repo.UpdateUserByID(ctx, id, user)
}

func (s *Service) DeleteUserByID(ctx context.Context, id string) error {
	return s.repo.DeleteUserByID(ctx, id)
}

func (s *Service) GetUserByID(ctx context.Context, id string) (User, error) {
	return s.repo.GetUserByID(ctx, id)
}

// generateID генерирует случайный ID для пользователя
func generateID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}