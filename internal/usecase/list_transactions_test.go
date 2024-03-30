package usecase_test

import (
	"errors"
	"testing"
	"time"

	"github.com/ortense/challenge-go/internal/payment/card"
	"github.com/ortense/challenge-go/internal/payment/method"
	"github.com/ortense/challenge-go/internal/repository/memory_repo"
	"github.com/ortense/challenge-go/internal/transaction"
	"github.com/ortense/challenge-go/internal/usecase"
	"github.com/ortense/challenge-go/pkg/currency"
	"github.com/ortense/challenge-go/pkg/id"
	"github.com/ortense/challenge-go/pkg/money"
	"github.com/stretchr/testify/assert"
)

func makeMockTransactionData() map[string]transaction.Transaction {
	first, _ := id.FromString("1")
	second, _ := id.FromString("2")

	return map[string]transaction.Transaction{
		first.String(): {
			Id:          first,
			Value:       money.New(10000, &currency.BRL),
			Description: "Red tshirt large",
			Method:      method.CreditCard,
			CreatedAt:   time.Date(2024, 1, 1, 10, 10, 0, 0, time.UTC),
			Card: card.PaymentCard{
				Number:     "1234",
				Holder:     "Jane Doe",
				Expiration: time.Date(2030, 1, 1, 10, 10, 0, 0, time.UTC),
				CVV:        "123",
			},
		},
		second.String(): {
			Id:          second,
			Value:       money.New(30000, &currency.BRL),
			Description: "Black dress small",
			Method:      method.DebitCard,
			CreatedAt:   time.Date(2024, 1, 2, 10, 10, 0, 0, time.UTC),
			Card: card.PaymentCard{
				Number:     "4567",
				Holder:     "Jane Doe",
				Expiration: time.Date(2030, 1, 1, 10, 10, 0, 0, time.UTC),
				CVV:        "999",
			},
		},
	}
}

func TestListTransactions(t *testing.T) {
	mockData := makeMockTransactionData()
	repo := memory_repo.NewTransactionMemoryRepo()
	repo.Data = mockData
	expected := []usecase.ListTransactionOutput{
		{
			Id:                 "1",
			Value:              10000,
			Currency:           "USD",
			DisplayValue:       "$100.00",
			Description:        "Red tshirt large",
			Method:             "CREDIT",
			CardNumber:         "1234",
			CardHolderName:     "Jane Doe",
			CardExpirationDate: "2030-01-01 10:10:00 +0000 UTC",
			CardCvv:            "123",
			CreatedAt:          "2024-01-01 10:10:00 +0000 UTC",
		},
		{
			Id:                 "2",
			Value:              30000,
			Currency:           "BRL",
			DisplayValue:       "R$ 300,00",
			Description:        "Black dress small",
			Method:             "DEBIT",
			CardNumber:         "4567",
			CardHolderName:     "Jane Doe",
			CardExpirationDate: "2030-01-01 10:10:00 +0000 UTC",
			CardCvv:            "999",
			CreatedAt:          "2024-01-02 10:10:00 +0000 UTC",
		},
	}

	output, err := usecase.ListTransactions(&repo)

	assert.Nil(t, err)
	assert.Equal(t, len(output), len(mockData))

	for i, out := range output {
		assert.Equal(t, expected[i].Id, out.Id)
		assert.Equal(t, expected[i].Description, out.Description)
		assert.Equal(t, expected[i].Method, out.Method)
		assert.Equal(t, expected[i].CardNumber, out.CardNumber)
		assert.Equal(t, expected[i].CardHolderName, out.CardHolderName)
		assert.Equal(t, expected[i].CardCvv, out.CardCvv)
		assert.Equal(t, expected[i].CreatedAt, out.CreatedAt)
	}
}

func TestListTransactions_Error(t *testing.T) {
	repo := memory_repo.NewTransactionMemoryRepo()
	repo.Err = errors.New("test repository error")

	output, err := usecase.ListTransactions(&repo)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "test repository error")
	assert.Equal(t, len(output), 0)
}
