package sql_repo

import (
	"github.com/ortense/challenge-go/internal/database"
	"github.com/ortense/challenge-go/internal/payable"
	"github.com/ortense/challenge-go/pkg/currency"
	"github.com/ortense/challenge-go/pkg/id"
	"github.com/ortense/challenge-go/pkg/money"
)

type Payable struct {
	connection database.Connection
}

func NewPayablenSqlRepo(connection database.Connection) Payable {
	return Payable{connection}
}

func (r *Payable) Save(p payable.Payable) error {
	err := r.connection.Open()

	if err != nil {
		return err
	}

	defer r.connection.Close()

	query := `
		INSERT INTO Payables(
			id,
			transaction_id, 
			status,
			currency,
			subtotal,
			total,
			discount, 
			created_at
		) VALUES(?, ?, ?, ?, ?, ?, ?, ?);`

	_, err = r.connection.Execute(
		query,
		p.Id.String(),
		p.TransactionId.String(),
		p.Status,
		p.Subtotal.Currency().Code(),
		p.Subtotal.Amount(),
		p.Total.Amount(),
		p.Discount,
		p.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Payable) List() ([]payable.Payable, error) {
	err := r.connection.Open()

	if err != nil {
		return nil, err
	}

	defer r.connection.Close()

	query := `
		SELECT 
			id,
			transaction_id,
			status,
			currency,
			subtotal,
			total,
			discount,
			created_at
		FROM Payables`

	rows, err := r.connection.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	list := []payable.Payable{}

	for rows.Next() {
		var p payable.Payable
		var pid, tid, currencyCode string
		var subtotal int
		var total int

		err := rows.Scan(
			&pid,
			&tid,
			&p.Status,
			&currencyCode,
			&subtotal,
			&total,
			&p.Discount,
			&p.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		p.Id, _ = id.FromString(pid)
		p.TransactionId, _ = id.FromString(tid)

		c, err := currency.FromCode(currencyCode)

		if err != nil {
			return nil, err
		}

		p.Subtotal = money.New(subtotal, c)
		p.Total = money.New(total, c)

		list = append(list, p)
	}

	return list, nil
}
