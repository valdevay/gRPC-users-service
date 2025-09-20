package user

import (
	"context"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context, limit, offset int) ([]User, error)
	CreateUser(ctx context.Context, user User) (User, error)
	UpdateUserByID(ctx context.Context, id string, user User) (User, error)
	DeleteUserByID(ctx context.Context, id string) error
	GetUserByID(ctx context.Context, id string) (User, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetAllUsers(ctx context.Context, limit, offset int) ([]User, error) {
	var users []User
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}

func (r *UserRepo) CreateUser(ctx context.Context, user User) (User, error) {
	err := r.db.WithContext(ctx).Create(&user).Error
	return user, err
}

func (r *UserRepo) UpdateUserByID(ctx context.Context, id string, user User) (User, error) {
	err := r.db.WithContext(ctx).Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return User{}, err
	}
	return r.GetUserByID(ctx, id)
}

func (r *UserRepo) DeleteUserByID(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&User{}, "id = ?", id).Error
}

func (r *UserRepo) GetUserByID(ctx context.Context, id string) (User, error) {
	var user User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return user, err
}