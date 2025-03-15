package router

import (
	api "example/GoApp/api"
	"fmt"
	"log"
	"net/http"

	//"github.com/gin-gonic/gin"

	"example/GoApp/repository"
	"example/GoApp/service"

	"gorm.io/gorm"
)

func AddAllAlbumAPIRoutes() {
	log.Println(" adding all album APIs to router")

	http.HandleFunc("/albums", api.GetAlbums) // add /album GET API to router
}

func AddUserAPIRoutes(db *gorm.DB) { // TODO:  clean this up and shift methods to service layer and api layer
	fmt.Println(" adding all user APIs to router")

	// Initialize repository
	userRepo := repository.NewUserRepository(db)

	// Initialize handlers
	userHandler := service.NewUserHandler(userRepo)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		case http.MethodGet:
			userHandler.ListUsers(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

}
