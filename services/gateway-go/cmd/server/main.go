package main

import (
	"log"

	"codesentinel/gateway/internal/config"
	"codesentinel/gateway/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()

	router := gin.Default()

	routes.RegisterRoutes(router)

	port := config.GetPort()

	log.Printf("Gateway running on port %s", port)

	router.Run(":" + port)
}