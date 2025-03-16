package router

import (
	api "example/GoApp/api"
	"fmt"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func AddAlbumAPIRoutes() {
	log.Println(" adding all album APIs to router")

	http.HandleFunc("/albums", api.GetAlbums) // add /album GET API to router
}

func AddUserAPIRoutes(db *gorm.DB) { // TODO:  clean this up and shift methods to service layer and api layer
	fmt.Println(" adding all user APIs to router")

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			api.CreateUser(db, w, r)
		case http.MethodGet:
			api.FetchUsers(db, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

}
