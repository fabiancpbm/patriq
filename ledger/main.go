package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"patriq.com.br/ledger/api"
	"patriq.com.br/ledger/db"
	"patriq.com.br/ledger/model"
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

func main() {
	database, err := db.Connect()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	transactionApi := &api.TransactionAPI{}
	transactionPersistence := &db.TransactionPersistence{}
	abstractApiConfig := &AbstractApiConfig[api.TransactionIn, model.Transaction, db.Transaction, api.TransactionOut]{
		Resource:   "transaction",
		Database:   database,
		Api: transactionApi,
		Persistence: transactionPersistence,
	}

	router.POST(
		"/transactions",
		func(c *gin.Context) {
			abstractApiConfig.Post(c)
		})
	router.Run(":8080")
}
