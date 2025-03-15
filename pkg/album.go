package api

import (
	"example/GoApp/pkg/model"
	"fmt"
	"net/http"
	// "github.com/gin-gonic/gin"
)

func GetAlbums(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, " non Get Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"message": "Hello, world!"}`, model.Albums)

}

func AddAllRoutesToRouter() {

}

// getAlbums responds with the list of all albums as JSON. via gin
//func GetAlbums(c *gin.Context) {
//	c.IndentedJSON(http.StatusOK, model.Albums) // using imported model class Albums
//}
