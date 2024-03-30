package sql_repo

import (
	"github.com/ortense/challenge-go/internal/database"
	"github.com/ortense/challenge-go/internal/payment/card"
	"github.com/ortense/challenge-go/internal/payment/method"
	"github.com/ortense/challenge-go/internal/transaction"
	"github.com/ortense/challenge-go/pkg/currency"
	"github.com/ortense/challenge-go/pkg/id"
	"github.com/ortense/challenge-go/pkg/money"
)

type Transaction struct {
	connection database.Connection
}

func NewTransactionSqlRepo(connection database.Connection) Transaction {
	return Transaction{connection}
}

func (r *Transaction) Save(t transaction.Transaction) error {
	err := r.connection.Open()

	if err != nil {
		return err
	}

	defer r.connection.Close()

	query := `
		INSERT INTO Transactions(
			id,
			description,
			value,
			currency,
			method,
			card_number,
			card_holder_name,
			card_expiration_date,
			card_cvv,
			created_at
		) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = r.connection.Execute(
		query,
		t.Id.String(),
		t.Description,
		t.Value.Amount(),
		t.Value.Currency().Code(),
		t.Method.Type,
		t.Card.Number,
		t.Card.Holder,
		t.Card.Expiration.Format("01/06"),
		t.Card.CVV,
		t.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Transaction) List() ([]transaction.Transaction, error) {
	err := r.connection.Open()

	if err != nil {
		return nil, err
	}

	defer r.connection.Close()

	query := `
		SELECT
			id,
			value,
			currency,
			description,
			method,
			card_number,
			card_holder_name,
			card_expiration_date,
			card_cvv,
			created_at
		FROM Transactions`

	rows, err := r.connection.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	list := []transaction.Transaction{}

	for rows.Next() {
		var t transaction.Transaction
		var value int
		var paymentType method.PaymentType
		var tid, currencyCode, cardNumber, cardHolderName, cardExpirationDate, cardCvv string

		err := rows.Scan(
			&tid,
			&value,
			&currencyCode,
			&t.Description,
			&paymentType,
			&cardNumber,
			&cardHolderName,
			&cardExpirationDate,
			&cardCvv,
			&t.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		t.Id, _ = id.FromString(tid)

		m, err := method.FromType(paymentType)

		if err != nil {
			return nil, err
		}

		t.Method = m

		c, err := currency.FromCode(currencyCode)

		if err != nil {
			return nil, err
		}

		t.Value = money.New(value, c)

		t.Card = card.New(
			cardNumber,
			cardHolderName,
			cardExpirationDate,
			cardCvv,
		)

		list = append(list, t)
	}

	return list, nil
}
