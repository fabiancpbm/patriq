package logic

import (
	"patriq.com.br/ledger/model"
)

type AccountLogic struct{}

func (accountLogic *AccountLogic) Validate(model *model.Account) (*model.Account, error) {
	return model, nil
}
