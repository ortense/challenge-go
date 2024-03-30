package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ortense/challenge-go/internal/usecase"
)

type WebserverHanlder func(Webserver) http.HandlerFunc

const (
	CreateTransaction = "POST /v1/transactions"
	ListTranctions    = "GET /v1/transactions"
	ListPayables      = "GET /v1/payables"
)

var router = map[string]WebserverHanlder{
	CreateTransaction: func(ws Webserver) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			decoder := json.NewDecoder(req.Body)

			var input usecase.CreateTransactionInput

			err := decoder.Decode(&input)

			if err != nil {
				ws.logger.Error(err.Error(), "entrypoint", CreateTransaction)
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "invalid request body")
				return
			}

			_, _, err = usecase.CreateTransaction(
				input,
				ws.transactionRepository,
				ws.payableRepository,
			)

			if err != nil {
				ws.logger.Error(err.Error(), "entrypoint", CreateTransaction)
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, err.Error())
				return
			}

			w.WriteHeader(http.StatusCreated)
		}
	},
	ListTranctions: func(ws Webserver) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			output, err := usecase.ListTransactions(ws.transactionRepository)

			if err != nil {
				ws.logger.Error(err.Error(), "entrypoint", ListTranctions)
				w.WriteHeader(http.StatusBadGateway)
				fmt.Fprintf(w, "unable to find transactions")
				return
			}

			response, err := json.Marshal(output)

			if err != nil {
				ws.logger.Error(err.Error(), "entrypoint", ListTranctions)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "internal error")
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("content-type", "application/json")
			fmt.Fprintf(w, string(response))
		}
	},
	ListPayables: func(ws Webserver) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			output, err := usecase.ListPayables(ws.payableRepository)

			if err != nil {
				ws.logger.Error(err.Error(), "entrypoint", ListPayables)
				w.WriteHeader(http.StatusBadGateway)
				fmt.Fprintf(w, "unable to find payables")
				return
			}

			response, err := json.Marshal(output)

			if err != nil {
				ws.logger.Error(err.Error(), "entrypoint", ListPayables)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "internal error")
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("content-type", "application/json")
			fmt.Fprintf(w, string(response))
		}
	},
}
