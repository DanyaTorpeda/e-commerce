package main

import (
	"e-commerce/internal/config"
	"e-commerce/internal/domains/server"
	"e-commerce/internal/handlers"
	"e-commerce/internal/repository"
	"e-commerce/internal/services"
	"e-commerce/pkg/logger"
)

func main() {
	logger := logger.InitLogger()
	cfg, err := config.LoadConfig(logger)
	if err != nil {
		logger.Fatalf("error occured loading config: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		DBName:   cfg.DB.DBName,
		SSLMode:  cfg.DB.SSLMode,
	})
	if err != nil {
		logger.Fatalf("error occured connecting to db: %s", err.Error())
	}
	defer db.Close()
	repos := repository.New(db, logger)
	servs := services.New(repos, logger)
	handler := handlers.New(servs, logger)
	srv := new(server.Server)
	if err := srv.Run(cfg.Server.Port, handler.InitRoutes()); err != nil {
		logger.Fatalf("error occured running server on port %s: %s", cfg.Server.Port, err.Error())
	}
}
