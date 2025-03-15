package repository

import (
	"example/GoApp/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository { // to create repo object
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *model.User) error { // create a new user
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUsers() ([]model.User, error) { // fetch all users
	var users []model.User

	result := r.db.Find(&users)
	return users, result.Error
}

func (r *UserRepository) GetUsersByEmails(emails []string) ([]model.User, error) { // fetch users by emails
	var users []model.User
	result := r.db.Where("email IN (?)", emails).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
