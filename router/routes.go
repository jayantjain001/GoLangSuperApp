package router

import (
	api "example/GoApp/api"
	"fmt"
	"net/http"
	//"github.com/gin-gonic/gin"
)

func AddAllAPIRoutes() {
	fmt.Println(" adding all APIs to router")

	http.HandleFunc("/albums", api.GetAlbums) // add /album GET API to router

	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)

	if err != nil { // if error !=null means some error is there
		fmt.Println("Error starting server:", err)
	}
}
