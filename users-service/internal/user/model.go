package user

// User represents a user in the system
type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Email     string `json:"email" gorm:"uniqueIndex;not null"`
	Password  string `json:"-" gorm:"not null"` // Hidden from JSON
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime"`
}
