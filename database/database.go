package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"

	"example/GoApp/model"
)

// DatabaseConfig holds the MySQL database configuration.
type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

// Database holds the GORM database instance.
type Database struct {
	*gorm.DB
}

// NewDatabase creates a new database connection.
func NewDatabase(config DatabaseConfig) (*Database, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(20)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &Database{DB: db}, nil
}

// GetUserByID retrieves a user by ID.
func (db *Database) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	result := db.First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// CreateUser creates a new user.
func (db *Database) CreateUser(user *model.User) error {
	result := db.Create(user)
	return result.Error
}
