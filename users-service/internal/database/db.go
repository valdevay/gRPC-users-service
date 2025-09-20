package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Загружаем переменные окружения из файла config.env
	if err := godotenv.Load("../config.env"); err != nil {
		log.Println("Warning: config.env file not found, using environment variables")
	}

	// Получаем значения из переменных окружения
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres123")
	dbname := getEnv("DB_NAME", "users_db")
	sslmode := getEnv("DB_SSLMODE", "disable")

	// Формируем строку подключения
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	log.Println("Database connected successfully")
	
	// Выполняем автоматические миграции
	if err := runMigrations(); err != nil {
		log.Fatalf("Could not run migrations: %v", err)
	}
}

// runMigrations выполняет автоматические миграции
func runMigrations() error {
	// Импортируем модель User
	// В реальном проекте лучше вынести это в отдельный пакет
	
	// Создаем таблицу users если её нет
	if err := DB.AutoMigrate(&User{}); err != nil {
		return fmt.Errorf("failed to migrate User model: %w", err)
	}
	
	log.Println("Migrations completed successfully")
	return nil
}

// User модель для миграций (временно здесь, лучше вынести в отдельный пакет)
type User struct {
	ID        string     `gorm:"primaryKey"`
	Email     string     `gorm:"uniqueIndex;not null"`
	Password  string     `gorm:"not null"`
	CreatedAt time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"index"`
}

// getEnv получает переменную окружения или возвращает значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}