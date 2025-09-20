package user
import (
	"time"
)

type User struct {
	ID        string     `json:"id" gorm:"primaryKey"`
	Email     string     `json:"email" gorm:"uniqueIndex;not null"`
	Password  string     `json:"password" gorm:"not null"`
	CreatedAt time.Time  `json:"-" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"-" gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"-" gorm:"index"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}