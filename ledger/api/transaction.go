package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"patriq.com.br/ledger/controller"
	"patriq.com.br/ledger/model"
)

type Transaction struct {
	SourceID string  `json:"sourceId"`
	TargetID string  `json:"targetId"`
	Date     string  `json:"date"`
	Amount   float32 `json:"amount"`
}

func dtoToModel(dto Transaction) (*model.Transaction, error) {
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
	return &model.Transaction {
		ID: uuid.New(),
		SourceID: sourceID,
		TargetID: targetID,
		Amount: dto.Amount,
		Date: date,
		EventDate: time.Now(),
	}, nil
}

func (dependences *Dependences) PostTransaction(c *gin.Context) {
	var transaction Transaction
	if err := c.BindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, `{"error": "could not read json for [POST:transaction]"}`)
		return
	}

	model, err := dtoToModel(transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, `{"error": "could not read json for [POST:transaction]"}`)
		return
	}

	saved, err := controller.CreateTransaction(dependences.Database, model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, `{"error": "` + err.Error() + `"}`)
		return
	}
	c.JSON(http.StatusProcessing, saved)
}
