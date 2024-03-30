package database

import (
	"database/sql"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

type Connection struct {
	logger         *slog.Logger
	sql            *sql.DB
	dataSourceName string
}

func New(dataSourceName string, logger *slog.Logger) Connection {
	return Connection{
		logger:         logger,
		dataSourceName: dataSourceName,
	}
}

func (connection *Connection) Open() error {
	db, err := sql.Open("sqlite3", connection.dataSourceName)

	if err != nil {
		connection.logger.Error(err.Error())
		return err
	}

	err = db.Ping()

	if err != nil {
		connection.logger.Error(err.Error())
		return err
	}

	connection.sql = db

	connection.logger.Info("Connected to database " + connection.dataSourceName)
	return nil
}

func (connection *Connection) Close() error {
	if connection.sql == nil {
		return nil
	}

	err := connection.sql.Close()

	if err != nil {
		connection.logger.Error(err.Error())
		return err
	}

	connection.logger.Info("Database disconected")
	return nil
}

func (connection *Connection) Seed() error {

	createTransactionTable := `
		CREATE TABLE IF NOT EXISTS Transactions (
			id TEXT PRIMARY KEY,
			description TEXT,
			value INTEGER,
			currency TEXT,
			method TEXT,
			card_number TEXT,
			card_holder_name TEXT,
			card_expiration_date TEXT,
			card_cvv TEXT,
			created_at DATETIME
		);
	`

	_, err := connection.sql.Exec(createTransactionTable)

	if err != nil {
		connection.logger.Error(err.Error(), "query", createTransactionTable)
		return err
	}

	createPayableTable := `
		CREATE TABLE IF NOT EXISTS Payables (
			id TEXT PRIMARY KEY,
			transaction_id TEXT,
			status TEXT,
			currency TEXT,
			subtotal INTEGER,
			total INTEGER,
			discount INTEGER,
			created_at DATETIME
		);
	`

	_, err = connection.sql.Exec(createPayableTable)

	if err != nil {
		connection.logger.Error(err.Error(), "query", createPayableTable)
		return err
	}

	return nil
}

func (connection *Connection) Execute(query string, args ...any) (sql.Result, error) {
	stmt, err := connection.sql.Prepare(query)

	if err != nil {
		connection.logger.Error(err.Error(), "query", query)
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(args...)

	if err != nil {
		connection.logger.Error(err.Error())
		return nil, err
	}

	return result, nil
}

func (connection *Connection) Query(query string) (*sql.Rows, error) {
	rows, err := connection.sql.Query(query)

	if err != nil {
		connection.logger.Error(err.Error(), "query", query)
		return nil, err
	}

	return rows, nil
}
