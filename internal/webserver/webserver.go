package webserver

import (
	"log/slog"
	"net/http"

	"github.com/ortense/challenge-go/internal/repository"
)

type Webserver struct {
	host                  string
	transactionRepository repository.TransactionRepository
	payableRepository     repository.PayableRepository
	mux                   *http.ServeMux
	logger                *slog.Logger
}

func New(
	host string,
	transactionRepository repository.TransactionRepository,
	payableRepository repository.PayableRepository,
	logger slog.Logger,
) Webserver {
	return Webserver{
		host:                  host,
		transactionRepository: transactionRepository,
		payableRepository:     payableRepository,
		mux:                   http.NewServeMux(),
		logger:                &logger,
	}
}

func (server Webserver) Launch() error {
	server.configure()

	server.logger.Info("launch server at " + server.host)

	return http.ListenAndServe(server.host, server.mux)
}

func (server Webserver) configure() {
	for pattern, hander := range router {
		server.logger.Info("add route " + pattern)
		server.mux.HandleFunc(pattern, hander(server))
	}
}
