package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Router
	r := gin.Default()
	// Initialize Routes
	initializeRoutes(r)
	// Run the server
	r.Run() // listen and serve on 0.0.0.0:8080
}
