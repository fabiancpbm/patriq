package adapter

import "patriq.com.br/bankimporter/model"

type TransactionConverter interface {
	Convert(line []string) (*model.Transaction, error)
}
