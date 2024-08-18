package port

import "patriq.com.br/bankimporter/model"

type TransactionImporter interface {
	ImportTransactions() ([]model.Transaction, error)
}
