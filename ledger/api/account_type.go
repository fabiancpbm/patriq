package api

import "patriq.com.br/ledger/model"

type AccoutTypeIn struct {
	ID model.AccountTypeID `json:"id"`
}

type AccountTypeOut struct {
	ID model.AccountTypeID `json:"id"`
}

type AccountTypeAPI struct {}

func (api *AccountTypeAPI) PostDtoToModel(dto *AccoutTypeIn) (*model.AccountType, error) {
	return &model.AccountType{
		ID: model.AccountTypeID(dto.ID),
	}, nil
}

func (api *AccountTypeAPI) PostModelToDto(model *model.AccountType) (*AccountTypeOut, error) {
	return &AccountTypeOut{
		ID: model.ID,
	}, nil
}
