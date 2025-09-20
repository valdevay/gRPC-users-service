package user

import (
	"errors"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(email, password string) (*User, error) {
	// Проверяем, существует ли пользователь с таким email
	existingUser, _ := s.repo.GetUserByEmail(email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	user := &User{
		Email:    email,
		Password: password, // В реальном проекте здесь должно быть хеширование пароля
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetUserByID(id uint) (*User, error) {
	return s.repo.GetUserByID(id)
}

func (s *Service) GetUserByEmail(email string) (*User, error) {
	return s.repo.GetUserByEmail(email)
}

func (s *Service) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *Service) UpdateUser(id uint, email, password string) (*User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// Проверяем, не занят ли новый email другим пользователем
	if email != user.Email {
		existingUser, _ := s.repo.GetUserByEmail(email)
		if existingUser != nil {
			return nil, errors.New("user with this email already exists")
		}
	}

	user.Email = email
	if password != "" {
		user.Password = password // В реальном проекте здесь должно быть хеширование пароля
	}

	if err := s.repo.UpdateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) DeleteUser(id uint) error {
	// Проверяем, существует ли пользователь
	_, err := s.repo.GetUserByID(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteUser(id)
}
