package transaction

import (
	"errors"
	"time"

	"github.com/ortense/challenge-go/internal/payment/card"
	"github.com/ortense/challenge-go/internal/payment/method"
	"github.com/ortense/challenge-go/pkg/currency"
	"github.com/ortense/challenge-go/pkg/id"
	"github.com/ortense/challenge-go/pkg/money"
)

type Transaction struct {
	Id          id.Id
	Value       money.Money
	Method      method.PaymentMethod
	Card        card.PaymentCard
	CreatedAt   time.Time
	Description string
}

func New(
	value uint,
	description string,
	paymentMethod method.PaymentType,
	cardNumber string,
	cardHolderName string,
	cardExpirationDate string,
	cardCvv string,
	currencyCode string,
) (Transaction, error) {

	pm, err := method.FromType(paymentMethod)

	if err != nil {
		return Transaction{}, err
	}

	c := card.New(cardNumber, cardHolderName, cardExpirationDate, cardCvv)

	if c.IsExpired() {
		return Transaction{}, errors.New("card expired")
	}

	cr, err := currency.FromCode(currencyCode)

	if err != nil {
		return Transaction{}, err
	}

	return Transaction{
		Id:          id.New(),
		Value:       money.New(int(value), cr),
		Description: description,
		Method:      pm,
		Card:        c,
		CreatedAt:   time.Now(),
	}, nil
}
