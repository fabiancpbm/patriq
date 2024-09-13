package api

import (
	"time"

	"github.com/google/uuid"
	"patriq.com.br/ledger/model"
)

type AccountIn struct {
	UserID    uuid.UUID           `json:"userId"`
	Name      string              `json:"name"`
	Type      model.AccountTypeID `json:"type"`
	CreatedAt time.Time           `json:"createdAt"`
}

type AccountOut struct {
	ID        uuid.UUID           `json:"Id"`
	UserID    uuid.UUID           `json:"userId"`
	Name      string              `json:"name"`
	Type      model.AccountTypeID `json:"type"`
	CreatedAt time.Time           `json:"createdAt"`
}

type AccountAPI struct{}

func (api *AccountAPI) PostDtoToModel(dto *AccountIn) (*model.Account, error) {
	return &model.Account{
		ID:        uuid.New(),
		UserID:    dto.UserID,
		Name:      dto.Name,
		Type:      dto.Type,
		CreatedAt: dto.CreatedAt,
	}, nil
}

func (api *AccountAPI) PostModelToDto(model *model.Account) (*AccountOut, error) {
	return &AccountOut{
		ID:        model.ID,
		UserID:    model.UserID,
		Name:      model.Name,
		Type:      model.Type,
		CreatedAt: model.CreatedAt,
	}, nil
}
