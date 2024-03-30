package money

import (
	"errors"

	"github.com/ortense/challenge-go/pkg/currency"
)

type Money struct {
	amount   int
	currency currency.Currency
}

func New(amount int, currency currency.Currency) Money {
	return Money{amount, currency}
}

func (money Money) Amount() int {
	return money.amount
}

func (money Money) Currency() currency.Currency {
	return money.currency
}

func (money Money) IsSameCurrency(other Money) bool {
	return money.currency.Code() == other.currency.Code()
}

func (money Money) Equal(other Money) bool {
	return money.IsSameCurrency(other) && money.amount == other.amount
}

func (money Money) Greater(other Money) (bool, error) {
	if money.IsSameCurrency(other) {
		return money.amount > other.amount, nil
	}

	return false, errors.New("comparison of different currencies")
}

func (money Money) Less(other Money) (bool, error) {
	if money.IsSameCurrency(other) {
		return money.amount < other.amount, nil
	}

	return false, errors.New("comparison of different currencies")
}

func (money Money) Format() string {
	return money.currency.Format(money.amount)
}
