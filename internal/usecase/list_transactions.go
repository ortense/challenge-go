package usecase

import "github.com/ortense/challenge-go/internal/transaction"

type ListTransactionOutput struct {
	Id                 string `json:"id"`
	Value              int    `json:"value"`
	Currency           string `json:"currency"`
	DisplayValue       string `json:"display_value"`
	Description        string `json:"description"`
	Method             string `json:"method"`
	CardNumber         string `json:"card_number"`
	CardHolderName     string `json:"card_holder_name"`
	CardExpirationDate string `json:"card_expiration_date"`
	CardCvv            string `json:"card_cvv"`
	CreatedAt          string `json:"created_at"`
}

type ListTransactionRepository interface {
	List() ([]transaction.Transaction, error)
}

func ListTransactions(repo ListTransactionRepository) ([]ListTransactionOutput, error) {

	transactions, err := repo.List()

	if err != nil {
		return nil, err
	}

	var output []ListTransactionOutput

	for _, t := range transactions {
		output = append(output, ListTransactionOutput{
			Id:                 t.Id.String(),
			Value:              t.Value.Amount(),
			Currency:           t.Value.Currency().Code(),
			DisplayValue:       t.Value.Format(),
			Description:        t.Description,
			Method:             string(t.Method.Type),
			CardNumber:         t.Card.Number,
			CardHolderName:     t.Card.Holder,
			CardExpirationDate: t.Card.Expiration.Format("01/06"),
			CardCvv:            t.Card.CVV,
			CreatedAt:          t.CreatedAt.UTC().String(),
		})
	}

	return output, nil
}
