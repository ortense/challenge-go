package method

import "errors"

type PaymentCode uint

const (
	CreditCode PaymentCode = iota + 100
	DebitCode
)

type PaymentType string

const (
	Credit PaymentType = "CREDIT"
	Debit  PaymentType = "DEBIT"
)

type PaymentTax uint

const (
	CreditTax PaymentTax = 4
	DebitTax  PaymentTax = 2
)

type PaymentMethod struct {
	Code PaymentCode
	Type PaymentType
	Tax  PaymentTax
}

var CreditCard = PaymentMethod{CreditCode, Credit, CreditTax}
var DebitCard = PaymentMethod{DebitCode, Debit, DebitTax}

var typeMap = map[PaymentType]PaymentMethod{
	Credit: CreditCard,
	Debit:  DebitCard,
}

func FromType(t PaymentType) (PaymentMethod, error) {
	m, ok := typeMap[t]

	if ok {
		return m, nil
	}

	return m, errors.New("unknown payment method")
}
