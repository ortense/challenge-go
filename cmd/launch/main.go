package main

import (
	"log/slog"
	"os"

	"github.com/ortense/challenge-go/config"
	"github.com/ortense/challenge-go/internal/database"
	"github.com/ortense/challenge-go/internal/repository/sql_repo"
	"github.com/ortense/challenge-go/internal/webserver"
)

func main() {
	c := config.GetConfig()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	connection := database.New(c.DataSourceName, logger)

	err := connection.Open()

	if err != nil {
		logger.Error("Unable to conect to database", err)
		os.Exit(1)
	}

	defer connection.Close()

	err = connection.Seed()

	if err != nil {
		logger.Error("Database seed error", err)
		os.Exit(1)
	}

	transactionRepo := sql_repo.NewTransactionSqlRepo(connection)
	payableRepo := sql_repo.NewPayablenSqlRepo(connection)

	server := webserver.New(
		c.Hostname,
		&transactionRepo,
		&payableRepo,
		*logger,
	)

	logger.Info("Using in sql repositories")

	err = server.Launch()

	if err != nil {
		logger.Error("Web server launch error", err)
		os.Exit(1)
	}
}
