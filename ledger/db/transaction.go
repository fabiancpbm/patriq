package db

import (
	"database/sql"

	"github.com/google/uuid"
	"patriq.com.br/ledger/model"
)

func findAllByUUID(database *sql.DB, columnName string, value uuid.UUID) ([]model.Transaction, error) {
	var transactions []model.Transaction

	rows, err := database.Query("SELECT * FROM transaction WHERE " + columnName + " = ?", value.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction model.Transaction
		err := rows.Scan(&transaction.ID, &transaction.SourceID, &transaction.TargetID, &transaction.EventDate, &transaction.Date, &transaction.Amount)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func FindById(database *sql.DB, idColumnName string, id uuid.UUID) (model.Transaction, error) {
	var transactions []model.Transaction

	transactions, err := findAllByUUID(database, idColumnName, id)

	return transactions[len(transactions)-1], err
}

func SaveTransaction(database *sql.DB, transaction model.Transaction) (*model.Transaction, error) {
	_, err := database.Exec(
		"INSERT INTO transaction (id, source_id, target_id, transaction_date, event_date, amount) VALUES (?, ?, ?, ?, ?, ?)",
		transaction.ID, transaction.SourceID, transaction.TargetID, transaction.Date, transaction.EventDate, transaction.Amount)
	if err != nil {
		return nil, err
	}
	savedTransaction, err := FindById(database, "id", transaction.ID)
	if err != nil {
		return nil, err
	}
	return &savedTransaction, nil
}
