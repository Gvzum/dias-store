package main

import (
	"fmt"
	"github.com/Gvzum/dias-store.git/api/routers"
	"github.com/Gvzum/dias-store.git/config"
	"log"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		fmt.Printf("Failed to load env file: %s", err)
	}

	router := routers.NewRouter()
	port := fmt.Sprintf(":%s", config.AppConfig.Server.SERVER_PORT)

	if err := router.Run(port); err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
