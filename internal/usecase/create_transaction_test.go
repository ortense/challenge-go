package usecase_test

import (
	"testing"

	"github.com/ortense/challenge-go/internal/payment/method"
	"github.com/ortense/challenge-go/internal/repository/memory_repo"
	"github.com/ortense/challenge-go/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {

	type TestCase struct {
		input    usecase.CreateTransactionInput
		total    int
		discount uint
		method   method.PaymentMethod
	}

	cases := map[string]TestCase{
		"create transaction and payable with debet card": {
			input: usecase.CreateTransactionInput{
				Value:              10000,
				Currency:           "BRL",
				Description:        "Test Transaction",
				CardNumber:         "1234567890123456",
				CardHolderName:     "John Doe",
				CardExpirationDate: "12/40",
				CardCvv:            "123",
				Method:             "DEBIT",
			},
			total:    9800,
			discount: 2,
			method:   method.DebitCard,
		},
		"create transaction and payable with credit card": {
			input: usecase.CreateTransactionInput{
				Value:              10000,
				Currency:           "USD",
				Description:        "Test Transaction",
				CardNumber:         "1234567890123456",
				CardHolderName:     "Jane Doe",
				CardExpirationDate: "12/40",
				CardCvv:            "123",
				Method:             "CREDIT",
			},
			total:    9600,
			discount: 4,
			method:   method.CreditCard,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			transactionRepo := memory_repo.NewTransactionMemoryRepo()
			payableRepo := memory_repo.NewPayablenMemoryRepo()
			actualTransaction, actualPayable, err := usecase.CreateTransaction(
				tc.input,
				&transactionRepo,
				&payableRepo,
			)

			assert.Nil(t, err)
			assert.Equal(t, actualPayable.Total.Amount(), tc.total)
			assert.Equal(t, actualPayable.Discount, tc.discount)
			assert.Equal(t, actualTransaction.Method.Type, tc.method.Type)
			assert.Equal(t, actualTransaction.Method.Code, tc.method.Code)
			assert.Equal(t, actualTransaction.Method.Tax, tc.method.Tax)
		})
	}

}

func TestCreateTransaction_Error(t *testing.T) {
	cases := map[string]usecase.CreateTransactionInput{
		"card expired": {
			Value:              10000,
			Description:        "Test Transaction",
			CardNumber:         "1234567890123456",
			CardHolderName:     "John Doe",
			CardExpirationDate: "12/00",
			CardCvv:            "123",
			Method:             "DEBIT",
		},
		"unknown payment method": {
			Value:              10000,
			Description:        "Test Transaction",
			CardNumber:         "1234567890123456",
			CardHolderName:     "Jane Doe",
			CardExpirationDate: "12/00",
			CardCvv:            "123",
			Method:             "PIX",
		},
	}

	for name, input := range cases {
		t.Run(name, func(t *testing.T) {
			transactionRepo := memory_repo.NewTransactionMemoryRepo()
			payableRepo := memory_repo.NewPayablenMemoryRepo()
			_, _, err := usecase.CreateTransaction(
				input,
				&transactionRepo,
				&payableRepo,
			)

			assert.NotNil(t, err)
			assert.Equal(t, err.Error(), name)
		})
	}
}
