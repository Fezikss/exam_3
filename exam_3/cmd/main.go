package main

import (
	"context"
	"exam3/api"
	_ "exam3/api/docs"
	"exam3/config"
	"exam3/pkg/logger"
	"exam3/service"
	"exam3/storage/postgres"
	"fmt"
)

func main() {

	cfg := config.Load()

	log := logger.New(cfg.ServiceName)

	store, err := postgres.New(context.Background(), cfg, log)
	if err != nil {
		log.Error("error while connecting to db: %v", logger.Error(err))
	}
	defer store.Close()

	services := service.New(store, log)

	server := api.New(services, log)

	if err := server.Run("localhost:8080"); err != nil {
		fmt.Printf("error while running server: %v\n", err)
	}
}
