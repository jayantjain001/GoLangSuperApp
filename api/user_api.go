package api

import (
	"net/http"

	"example/GoApp/repository"
	"example/GoApp/service"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// Initialize repository
	userRepo := repository.NewUserRepository(db)

	// Initialize handlers
	userHandler := service.NewUserHandler(userRepo)

	userHandler.CreateUser(w, r)

}

func FetchUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	// initialize repository
	userRepo := repository.NewUserRepository(db)

	// Initialize handlers
	userHandler := service.NewUserHandler(userRepo)

	userHandler.ListUsers(w, r)

}
