package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go-project-template/internal/config"
	"go-project-template/internal/logger"
	"go-project-template/internal/migrations"
	"go-project-template/internal/repository"

	"go.uber.org/zap"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT)
	defer stop()

	logger := logger.Logger

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Panic("Error in loading config",
			zap.Error(err),
		)
	}
	logger.Info("Config loaded")
	err = migrations.RunMigrations(ctx, cfg.GetDBConnString(), cfg.DB_NAME, cfg.MIGRATIONS_SOURCE)
	if err != nil {
		logger.Panic("Error in running migrations",
			zap.Error(err),
		)
	}
	logger.Info("Migrations completed")
	_, err = repository.NewRepository(ctx, cfg.GetDBConnString())
	if err != nil {
		logger.Panic("Error in creating dao object",
			zap.Error(err),
		)
	}
	logger.Info("Dao object created")
}
