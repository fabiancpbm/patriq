package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"patriq.com.br/ledger/model"
)

type Transaction struct {
	ID         uuid.UUID `db:"transaction__id"`
	SourceID   uuid.UUID `db:"transaction__source_id"`
	TargetID   uuid.UUID `db:"transaction__target_id"`
	Date       time.Time `db:"transaction__date"`
	Amount     float32   `db:"transaction__amount"`
	EventDate  time.Time `db:"transaction__event_date"`
	CategoryID uuid.UUID `db:"category__id"`
}

type TransactionPersistence struct {
}

func (transactionPersistence *TransactionPersistence) ModelToEntity(model *model.Transaction) (*Transaction, error) {
	return &Transaction{
		ID:         model.ID,
		SourceID:   model.SourceID,
		TargetID:   model.TargetID,
		Date:       model.Date,
		Amount:     model.Amount,
		EventDate:  model.EventDate,
		CategoryID: model.CategoryID,
	}, nil
}

func (transactionPersistence *TransactionPersistence) EntityToModel(entity *Transaction) (*model.Transaction, error) {
	return &model.Transaction{
		ID:         entity.ID,
		SourceID:   entity.SourceID,
		TargetID:   entity.TargetID,
		Date:       entity.Date,
		Amount:     entity.Amount,
		EventDate:  entity.EventDate,
		CategoryID: entity.CategoryID,
	}, nil
}

func (transactionPersistence *TransactionPersistence) Save(database *sql.DB, entity *Transaction) (*Transaction, error) {
	_, err := database.Exec(`
		INSERT INTO ledger__transaction(
			transaction__id,
			transaction__source_id,
			transaction__target_id,
			transaction__date,
			transaction__event_date,
			transaction__amount,
			category__id)
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		entity.ID, entity.SourceID, entity.TargetID, entity.Date, entity.EventDate, entity.Amount, entity.CategoryID)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func findAllByUUID(database *sql.DB, columnName string, value uuid.UUID) ([]model.Transaction, error) {
	var transactions []model.Transaction

	rows, err := database.Query("SELECT * FROM transaction WHERE "+columnName+" = ?", value.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction model.Transaction
		err := rows.Scan(&transaction.ID, &transaction.SourceID, &transaction.TargetID, &transaction.EventDate, &transaction.Date, &transaction.Amount)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func FindById(database *sql.DB, idColumnName string, id uuid.UUID) (model.Transaction, error) {
	var transactions []model.Transaction

	transactions, err := findAllByUUID(database, idColumnName, id)

	return transactions[len(transactions)-1], err
}

func SaveTransaction(database *sql.DB, transaction model.Transaction) (*model.Transaction, error) {
	_, err := database.Exec(
		"INSERT INTO transaction (id, source_id, target_id, transaction_date, event_date, amount) VALUES (?, ?, ?, ?, ?, ?)",
		transaction.ID, transaction.SourceID, transaction.TargetID, transaction.Date, transaction.EventDate, transaction.Amount)
	if err != nil {
		return nil, err
	}
	savedTransaction, err := FindById(database, "id", transaction.ID)
	if err != nil {
		return nil, err
	}
	return &savedTransaction, nil
}
