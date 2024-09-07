package api

import (
	"time"

	"github.com/google/uuid"
	"patriq.com.br/ledger/model"
)

type Account struct {
	UserID    uuid.UUID         `json:"userId"`
	Name      string            `json:"name"`
	Type      model.AccountType `json:"type"`
	CreatedAt time.Time         `json:"createdAt"`
}

// func accountDtoToModel(dto Account) (*model.Account, error) {
// 	return &model.Account{
// 		ID:        uuid.New(),
// 		UserID:    dto.UserID,
// 		Name:      dto.Name,
// 		Type:      dto.Type,
// 		CreatedAt: dto.CreatedAt,
// 	}, nil
// }
