package controller

import (
	"database/sql"

	"patriq.com.br/ledger/db"
	"patriq.com.br/ledger/model"
)

func CreateTransaction(database *sql.DB, transaction *model.Transaction) (*model.Transaction, error) {
    saved, err := db.SaveTransaction(database, *transaction)
	if err != nil {
		return nil, err
	}

	return saved, nil
}
