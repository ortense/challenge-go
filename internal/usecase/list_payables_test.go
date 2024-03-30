package usecase_test

import (
	"errors"
	"testing"
	"time"

	"github.com/ortense/challenge-go/internal/payable"
	"github.com/ortense/challenge-go/internal/repository/memory_repo"
	"github.com/ortense/challenge-go/internal/usecase"
	"github.com/ortense/challenge-go/pkg/currency"
	"github.com/ortense/challenge-go/pkg/id"
	"github.com/ortense/challenge-go/pkg/money"
	"github.com/stretchr/testify/assert"
)

func makeMockPayableData() map[string]payable.Payable {
	fist, _ := id.FromString("1")
	second, _ := id.FromString("2")
	transactionId, _ := id.FromString("3")

	return map[string]payable.Payable{
		fist.String(): {
			Id:            fist,
			TransactionId: transactionId,
			Subtotal:      money.New(10000, &currency.BRL),
			Discount:      2,
			Total:         money.New(9800, &currency.BRL),
			Status:        payable.PayableStatusPaid,
			CreatedAt:     time.Date(2024, 1, 1, 10, 10, 0, 0, time.UTC),
		},
		second.String(): {
			Id:            second,
			TransactionId: transactionId,
			Subtotal:      money.New(10000, &currency.USD),
			Discount:      4,
			Total:         money.New(9600, &currency.USD),
			Status:        payable.PayableStatusWaitingFunds,
			CreatedAt:     time.Date(2024, 2, 2, 2, 2, 0, 0, time.UTC),
		},
	}
}

func TestListPayables(t *testing.T) {
	mockData := makeMockPayableData()
	repo := memory_repo.NewPayablenMemoryRepo()
	repo.Data = mockData
	expected := []usecase.ListPayableOutput{
		{
			Id:            "1",
			TransactionId: "3",
			Status:        "paid",
			Subtotal:      10000,
			Discount:      2,
			Total:         9800,
			CreatedAt:     "2024-01-01 10:10:00 +0000 UTC",
		},
		{
			Id:            "2",
			TransactionId: "3",
			Status:        "waiting_funds",
			Subtotal:      10000,
			Discount:      4,
			Total:         9600,
			CreatedAt:     "2024-02-02 02:02:00 +0000 UTC",
		},
	}

	output, err := usecase.ListPayables(&repo)

	assert.Nil(t, err)
	assert.Equal(t, len(output), len(mockData))
	for i, out := range output {
		assert.Equal(t, expected[i].Id, out.Id)
		assert.Equal(t, expected[i].TransactionId, out.TransactionId)
		assert.Equal(t, expected[i].Status, out.Status)
		assert.Equal(t, expected[i].Subtotal, out.Subtotal)
		assert.Equal(t, expected[i].Discount, out.Discount)
		assert.Equal(t, expected[i].Total, out.Total)
		assert.Equal(t, expected[i].CreatedAt, out.CreatedAt)
	}
}

func TestListPayables_Error(t *testing.T) {
	repo := memory_repo.NewPayablenMemoryRepo()
	repo.Err = errors.New("test repository error")

	output, err := usecase.ListPayables(&repo)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "test repository error")
	assert.Equal(t, len(output), 0)
}
