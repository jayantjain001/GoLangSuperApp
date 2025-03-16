package main

import (
	"example/GoApp/config"
	"example/GoApp/database"
	"example/GoApp/router"

	"fmt"
	"log"
	"net/http"
	"runtime"
)

func main() {
	fmt.Println("Starting Go App !")
	fmt.Println(" machine cores : ", runtime.NumCPU())

	// Load configuration
	cfg, err := config.LoadDatabaseConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Initialize database
	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Register routes
	router.AddAlbumAPIRoutes()  //  registering all apis to router
	router.AddUserAPIRoutes(db) // registering all user apis to router

	// Start http server
	log.Println("Server starting on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
