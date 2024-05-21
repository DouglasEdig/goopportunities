package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initialize Router
	router := gin.Default()

	//Initializes Routes
	initializeRoutes(router)

	//run the server
	router.Run(":8080")
}
