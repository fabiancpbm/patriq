package logic

import (
	"errors"

	"patriq.com.br/ledger/model"
)

type AccountTypeLogic struct{}

func (accountTypeLogic *AccountTypeLogic) Validate(model *model.AccountType) (*model.AccountType, error) {
	if isAccountTypeID(string(model.ID)) {
		return model, nil
	}
	return nil, errors.New("invalid account type")
}

func isAccountTypeID(id string) bool {
	switch model.AccountTypeID(id) {
	case model.Asset, model.Liability, model.Revenue, model.Expense:
		return true
	default:
		return false
	}
}
