package main

import (
	"log/slog"
	"os"

	"github.com/ortense/challenge-go/config"
	"github.com/ortense/challenge-go/internal/repository/memory_repo"
	"github.com/ortense/challenge-go/internal/webserver"
)

func main() {
	c := config.GetConfig()
	transactionRepo := memory_repo.NewTransactionMemoryRepo()
	payableRepo := memory_repo.NewPayablenMemoryRepo()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	server := webserver.New(
		c.Hostname,
		&transactionRepo,
		&payableRepo,
		*logger,
	)

	logger.Info("Using in memory repositories")

	err := server.Launch()

	if err != nil {
		logger.Error("Web server launch error", err)
		os.Exit(1)
	}
}
