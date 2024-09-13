package logic

import (
	"patriq.com.br/ledger/model"
)

type TransactionLogic struct{}

func (transactionLogic *TransactionLogic) Validate(model *model.Transaction) (*model.Transaction, error) {
	return model, nil
}
