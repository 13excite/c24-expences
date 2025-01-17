// package models contains the models for the application and the database connection.
package models

import (
	"context"
	"database/sql"
	"time"
)

// DBModel is the type for db connection values
type DBModel struct {
	DB *sql.DB
}

// Models is the wrapper for all models
type Models struct {
	DB DBModel
}

// NewModels returns a model type with database connection pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// Transaction struct that holds the transaction data
type Transaction struct {
	TransactionType string
	Date            string
	Amount          float64
	Recipient       string
	Usage           string
	Category        string
	Subcategory     string
}

// InsertTransaction inserts a new txn and returns its id
func (m *DBModel) InsertTransaction(txn Transaction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		INSERT INTO transactions
			(kind, date, recipient,
			 amount, primary_class, secondary_class)
		VALUES (?, ?, ?, ?, ?, ?)
		`
	_, err := m.DB.ExecContext(ctx, stmt,
		txn.TransactionType, txn.Date, txn.Recipient,
		txn.Amount, txn.Category, txn.Subcategory,
	)

	if err != nil {
		return err
	}

	return nil
}
