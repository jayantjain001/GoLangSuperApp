package main

import (
	"example/GoApp/config"
	"example/GoApp/database"
	"net/http"

	//"example/GoApp/repository"
	"example/GoApp/router"
	///	"example/GoApp/service"

	"fmt"
	"log"
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

	//	config.LoadDatabaseConfig()

	// Register routes

	router.AddAllAlbumAPIRoutes() //  registering all apis to router
	router.AddUserAPIRoutes(db)   // registering all user apis to router

	// Start server
	log.Println("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

//gin.SetMode(gin.ReleaseMode)         // running release mode
///router := gin.Default()              // get gin router
//router.GET("/albums", api.GetAlbums) // set api and its caller method via gin
//router.Run("localhost:8080")         // run gin router
