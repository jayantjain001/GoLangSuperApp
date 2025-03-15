package main

import (
	"example/GoApp/config"
	"example/GoApp/router"
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Starting Go App !")
	fmt.Println(" machine cores : ", runtime.NumCPU())

	config.LoadDatabaseConfig() // load database config

	router.AddAllAPIRoutes() //  registering all apis

}

//gin.SetMode(gin.ReleaseMode)         // running release mode
///router := gin.Default()              // get gin router
//router.GET("/albums", api.GetAlbums) // set api and its caller method via gin
//router.Run("localhost:8080")         // run gin router
