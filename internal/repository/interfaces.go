package repository

import (
	"github.com/ortense/challenge-go/internal/payable"
	"github.com/ortense/challenge-go/internal/transaction"
)

type TransactionRepository interface {
	Save(transaction.Transaction) error
	List() ([]transaction.Transaction, error)
}

type PayableRepository interface {
	Save(payable.Payable) error
	List() ([]payable.Payable, error)
}
