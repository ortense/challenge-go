package payable

import (
	"errors"
	"time"

	"github.com/ortense/challenge-go/internal/payment/method"
	"github.com/ortense/challenge-go/internal/transaction"
	"github.com/ortense/challenge-go/pkg/id"
	"github.com/ortense/challenge-go/pkg/money"
)

type PayableStatus string

const (
	PayableStatusPaid         PayableStatus = "paid"
	PayableStatusWaitingFunds PayableStatus = "waiting_funds"
)

var paymentTypeStatus = map[method.PaymentType]PayableStatus{
	method.Credit: PayableStatusWaitingFunds,
	method.Debit:  PayableStatusPaid,
}

type Payable struct {
	Id            id.Id
	TransactionId id.Id
	Subtotal      money.Money
	Discount      uint
	Total         money.Money
	Status        PayableStatus
	CreatedAt     time.Time
}

func FromTransaction(t transaction.Transaction) (Payable, error) {

	status, ok := paymentTypeStatus[t.Method.Type]

	if !ok {
		return Payable{}, errors.New("unknown payment method")
	}

	discount, total := applyDiscount(t)

	return Payable{
		Id:            id.New(),
		TransactionId: t.Id,
		Subtotal:      t.Value,
		Status:        status,
		Discount:      discount,
		Total:         total,
		CreatedAt:     time.Now(),
	}, nil
}

func applyDiscount(t transaction.Transaction) (discount uint, total money.Money) {
	discount = uint(t.Method.Tax)
	subtotal := t.Value.Amount()
	tax := (subtotal / 100) * int(t.Method.Tax)
	amount := subtotal - tax

	total = money.New(amount, t.Value.Currency())

	return discount, total
}
