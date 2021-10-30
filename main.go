package main

import (
	"github.com/alexdang1993374/employee-api/config"
	"github.com/alexdang1993374/employee-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()

	router := gin.Default()

	routes.Routes(router)

	router.Run(":5000")
}
