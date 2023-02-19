package main

import (
	"fmt"
	"log"

	"github.com/Gvzum/dias-store.git/cmd/server/config"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Failed to load env file: %s", err)
	}

	router := gin.Default()

	port := config.AppConfig.Server.SERVER_PORT

	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
