// Package models contains the models for the application and the database connection.
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

// Model is the wrapper for all models
type Model struct {
	DB DBModel
}

// NewModel returns a model type with database connection pool
func NewModel(db *sql.DB) Model {
	return Model{
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

// SHAFile struct that holds the path and sha of a file for preventing duplicate uploads
type SHAFile struct {
	Path   string
	SHA256 string
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

// GetSHAFiles retrieves all SHA files from the database
func (m *DBModel) GetSHAFiles() ([]SHAFile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT path, sha256 FROM file_hashes`

	rows, err := m.DB.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shaFiles []SHAFile
	for rows.Next() {
		var shaFile SHAFile
		err := rows.Scan(&shaFile.Path, &shaFile.SHA256)
		if err != nil {
			return nil, err
		}
		shaFiles = append(shaFiles, shaFile)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return shaFiles, nil
}

// InsertSHAFile inserts a new SHA file into the database
func (m *DBModel) InsertSHAFile(shaFile SHAFile) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		INSERT INTO file_hashes
			(path, sha256)
		VALUES (?, ?)
		`
	_, err := m.DB.ExecContext(ctx, stmt,
		shaFile.Path, shaFile.SHA256,
	)

	if err != nil {
		return err
	}

	return nil
}
