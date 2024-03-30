package usecase

import (
	"github.com/ortense/challenge-go/internal/payable"
	"github.com/ortense/challenge-go/internal/payment/method"
	"github.com/ortense/challenge-go/internal/transaction"
)

type CreateTransactionInput struct {
	Value              uint               `json:"value"`
	Currency           string             `json:"currency"`
	Description        string             `json:"description"`
	CardNumber         string             `json:"card_number"`
	CardHolderName     string             `json:"card_holder_name"`
	CardExpirationDate string             `json:"card_expiration_date"`
	CardCvv            string             `json:"card_cvv"`
	Method             method.PaymentType `json:"method"`
}

type CreateTransactionRepository interface {
	Save(transaction.Transaction) error
}

type CreatePayableRepository interface {
	Save(payable.Payable) error
}

func CreateTransaction(
	input CreateTransactionInput,
	transactionRepo CreateTransactionRepository,
	paybableRepo CreatePayableRepository,
) (*transaction.Transaction, *payable.Payable, error) {

	t, err := transaction.New(
		input.Value,
		input.Description,
		input.Method,
		input.CardNumber,
		input.CardHolderName,
		input.CardExpirationDate,
		input.CardCvv,
		input.Currency,
	)

	if err != nil {
		return &t, nil, err
	}

	p, err := payable.FromTransaction(t)

	if err != nil {
		return nil, nil, err
	}

	err = transactionRepo.Save(t)

	if err != nil {
		return nil, nil, err
	}

	err = paybableRepo.Save(p)

	if err != nil {
		return nil, nil, err
	}

	return &t, &p, nil
}
