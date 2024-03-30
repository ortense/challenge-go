package usecase

import (
	"github.com/ortense/challenge-go/internal/payable"
)

type ListPayableOutput struct {
	Id              string `json:"id"`
	TransactionId   string `json:"transaction_id"`
	Status          string `json:"status"`
	Currency        string `json:"currency"`
	Subtotal        int    `json:"subtotal"`
	DisplaySubtotal string `json:"display_subtotal"`
	Total           int    `json:"total"`
	DisplayTotal    string `json:"display_total"`
	Discount        uint   `json:"discount"`
	CreatedAt       string `json:"created_at"`
}

type ListPayableRepository interface {
	List() ([]payable.Payable, error)
}

func ListPayables(repo ListPayableRepository) ([]ListPayableOutput, error) {

	payables, err := repo.List()

	if err != nil {
		return nil, err
	}

	var output []ListPayableOutput

	for _, p := range payables {
		output = append(output, ListPayableOutput{
			Id:              p.Id.String(),
			TransactionId:   p.TransactionId.String(),
			Status:          string(p.Status),
			Currency:        p.Total.Currency().Code(),
			Subtotal:        p.Subtotal.Amount(),
			DisplaySubtotal: p.Subtotal.Format(),
			Total:           p.Total.Amount(),
			DisplayTotal:    p.Total.Format(),
			Discount:        p.Discount,
			CreatedAt:       p.CreatedAt.UTC().String(),
		})
	}

	return output, nil
}
