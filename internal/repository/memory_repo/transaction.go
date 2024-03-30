package memory_repo

import (
	"errors"

	"github.com/ortense/challenge-go/internal/transaction"
)

type Transaction struct {
	Data map[string]transaction.Transaction
	Err  error
}

func NewTransactionMemoryRepo() Transaction {
	return Transaction{
		Data: map[string]transaction.Transaction{},
		Err:  nil,
	}
}

func (r *Transaction) Save(t transaction.Transaction) error {
	id := t.Id.String()

	if r.Err != nil {
		return r.Err
	}

	_, exist := r.Data[id]

	if exist {
		return errors.New("transaction id conflict")
	}

	r.Data[id] = t

	return nil
}

func (r *Transaction) List() ([]transaction.Transaction, error) {
	list := []transaction.Transaction{}

	if r.Err != nil {
		return list, r.Err
	}

	for _, t := range r.Data {
		list = append(list, t)
	}

	return list, nil
}
