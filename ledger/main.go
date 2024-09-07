package main

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"patriq.com.br/ledger/api"
	"patriq.com.br/ledger/db"
)

type DTOTransaction struct {
	SourceID string  `json:"sourceId"`
	TargetID string  `json:"targetId"`
	Date     string  `json:"date"`
	Amount   float32 `json:"amount"`
}
type MTransaction struct {
	ID        uuid.UUID `db:"id"`
	SourceID  uuid.UUID `db:"source_id"`
	TargetID  uuid.UUID `db:"target_id"`
	Date      time.Time `db:"transaction_date"`
	Amount    float32   `db:"amount"`
	EventDate time.Time `db:"event_date"`
}

type TransactionAPI struct {
	api.AbstractAPI[DTOTransaction, MTransaction]
}

func (api *TransactionAPI) DTOToModel(dto DTOTransaction) (*MTransaction, error) {
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
	return &MTransaction{
		ID:        uuid.New(),
		SourceID:  sourceID,
		TargetID:  targetID,
		Amount:    dto.Amount,
		Date:      date,
		EventDate: time.Now(),
	}, nil
}

func (api *TransactionAPI) Save(database *sql.DB, transaction MTransaction) (*MTransaction, error) {
	_, err := database.Exec(
		"INSERT INTO transaction (id, source_id, target_id, transaction_date, event_date, amount) VALUES (?, ?, ?, ?, ?, ?)",
		transaction.ID, transaction.SourceID, transaction.TargetID, transaction.Date, transaction.EventDate, transaction.Amount)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func main() {
	database, err := db.Connect()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	
	transactionApi := &TransactionAPI{
		AbstractAPI: api.AbstractAPI[DTOTransaction, MTransaction]{
			Resource: "transaction",
			Database: database,
		},
	}

	router.POST(
		"/transactions", 
		func(c *gin.Context) {
			transactionApi.Post(c, transactionApi)
		})
	router.Run(":8080")
}
