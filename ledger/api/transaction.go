package api

import (
	"time"

	"github.com/google/uuid"
	"patriq.com.br/ledger/model"
)

type TransactionIn struct {
	SourceID   uuid.UUID `json:"sourceId"`
	TargetID   uuid.UUID `json:"targetId"`
	Date       time.Time `json:"date"`
	Amount     float32   `json:"amount"`
	CategoryID uuid.UUID `json:"categoryId"`
}

type TransactionOut struct {
	ID         uuid.UUID `json:"id"`
	SourceID   uuid.UUID `json:"sourceId"`
	TargetID   uuid.UUID `json:"targetId"`
	Date       time.Time `json:"date"`
	Amount     float32   `json:"amount"`
	EventDate  time.Time `json:"eventDate"`
	CategoryID uuid.UUID `json:"categoryId"`
}

type TransactionAPI struct{}

func (api *TransactionAPI) PostDtoToModel(dto *TransactionIn) (*model.Transaction, error) {
	return &model.Transaction{
		ID:         uuid.New(),
		SourceID:   dto.SourceID,
		TargetID:   dto.TargetID,
		Amount:     dto.Amount,
		Date:       dto.Date,
		EventDate:  time.Now(),
		CategoryID: dto.CategoryID,
	}, nil
}

func (api *TransactionAPI) PostModelToDto(model *model.Transaction) (*TransactionOut, error) {
	return &TransactionOut{
		ID:         model.ID,
		SourceID:   model.SourceID,
		TargetID:   model.TargetID,
		Amount:     model.Amount,
		Date:       model.Date,
		EventDate:  model.EventDate,
		CategoryID: model.CategoryID,
	}, nil
}
