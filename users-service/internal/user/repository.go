package user

import (
	"database/sql"
	"errors"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(user *User) error {
	query := `INSERT INTO users (email, password, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`
	now := time.Now().Unix()
	err := r.db.QueryRow(query, user.Email, user.Password, now, now).Scan(&user.ID)
	if err != nil {
		return err
	}
	user.CreatedAt = now
	user.UpdatedAt = now
	return nil
}

func (r *Repository) GetUserByID(id uint) (*User, error) {
	query := `SELECT id, email, password, created_at, updated_at FROM users WHERE id = $1`
	user := &User{}
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (r *Repository) GetUserByEmail(email string) (*User, error) {
	query := `SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1`
	user := &User{}
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (r *Repository) GetAllUsers() ([]User, error) {
	query := `SELECT id, email, password, created_at, updated_at FROM users ORDER BY id`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *Repository) UpdateUser(user *User) error {
	query := `UPDATE users SET email = $1, password = $2, updated_at = $3 WHERE id = $4`
	now := time.Now().Unix()
	result, err := r.db.Exec(query, user.Email, user.Password, now, user.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	user.UpdatedAt = now
	return nil
}

func (r *Repository) DeleteUser(id uint) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
