package main

import (
	"github.com/alexdang1993374/employee-api/config"
	"github.com/alexdang1993374/employee-api/routes"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./", true)))

	// Setup route group for the API
	routes.Routes(router)

	// Start and run the server
	router.Run(":5000")
}
