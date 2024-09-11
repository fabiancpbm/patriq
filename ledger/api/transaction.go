package api

import (
	"time"

	"github.com/google/uuid"
	"patriq.com.br/ledger/model"
)

type TransactionIn struct {
	SourceID string  `json:"sourceId"`
	TargetID string  `json:"targetId"`
	Date     string  `json:"date"`
	Amount   float32 `json:"amount"`
}

type TransactionOut struct {
	ID       string  `json:"id"`
	SourceID string  `json:"sourceId"`
	TargetID string  `json:"targetId"`
	Date     string  `json:"date"`
	Amount   float32 `json:"amount"`
}

type TransactionAPI struct {}

func (api *TransactionAPI) PostDtoToModel(dto *TransactionIn) (*model.Transaction, error) {
	sourceID, err := uuid.Parse(dto.SourceID)
	if err != nil {
		return nil, err
	}
	targetID, err := uuid.Parse(dto.TargetID)
	if err != nil {
		return nil, err
	}
	date, err := time.Parse(time.RFC3339, dto.Date)
	if err != nil {
		return nil, err
	}
	return &model.Transaction{
		ID:        uuid.New(),
		SourceID:  sourceID,
		TargetID:  targetID,
		Amount:    dto.Amount,
		Date:      date,
		EventDate: time.Now(),
	}, nil
}

func (api *TransactionAPI) PostModelToDto(model *model.Transaction) (*TransactionOut, error) {
	return nil, nil
}
