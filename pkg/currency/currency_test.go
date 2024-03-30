package currency_test

import (
	"testing"

	"github.com/ortense/challenge-go/pkg/currency"
	"github.com/stretchr/testify/assert"
)

func TestCurrency(t *testing.T) {
	type TestCase struct {
		amount   int
		code     string
		symbol   string
		formated string
		currency currency.Currency
	}

	cases := map[string]TestCase{
		"format 10 in BRL": {
			amount:   10,
			code:     "BRL",
			symbol:   "R$",
			formated: "R$ 0,10",
			currency: &currency.BRL,
		},
		"format 10000 in BRL": {
			amount:   10000,
			code:     "BRL",
			symbol:   "R$",
			formated: "R$ 100,00",
			currency: &currency.BRL,
		},
		"format 10 in USD": {
			amount:   10,
			code:     "USD",
			symbol:   "$",
			formated: "$0.10",
			currency: &currency.USD,
		},
		"format 10000 in USD": {
			amount:   10000,
			code:     "USD",
			symbol:   "$",
			formated: "$100.00",
			currency: &currency.USD,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.currency.Code(), tc.code)
			assert.Equal(t, tc.currency.Symbol(), tc.symbol)
			assert.Equal(t, tc.currency.Format(tc.amount), tc.formated)
		})
	}
}
