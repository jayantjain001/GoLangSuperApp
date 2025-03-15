package api

import (
	"encoding/json"
	"example/GoApp/model"
	"log"
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

	log.Println(" album get api called ...")
	json.NewEncoder(w).Encode(model.Albums)
	//fmt.Fprint(w, `{"message": "Hello, world!"}`, model.Albums)

}
