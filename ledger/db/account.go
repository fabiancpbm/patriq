package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"patriq.com.br/ledger/model"
)

type Account struct {
	ID        uuid.UUID           `db:"account__id"`
	UserID    uuid.UUID           `db:"user__id"`
	Name      string              `db:"account__name"`
	Type      model.AccountTypeID `db:"account_type__id"`
	CreatedAt time.Time           `db:"account__created_at"`
}

type AccountPersistence struct {
}

func (accountPersistence *AccountPersistence) ModelToEntity(model *model.Account) (*Account, error) {
	return &Account{
		ID:        model.ID,
		UserID:    model.UserID,
		Name:      model.Name,
		Type:      model.Type,
		CreatedAt: model.CreatedAt,
	}, nil
}

func (accountPersistence *AccountPersistence) EntityToModel(entity *Account) (*model.Account, error) {
	return &model.Account{
		ID:        entity.ID,
		UserID:    entity.UserID,
		Name:      entity.Name,
		Type:      entity.Type,
		CreatedAt: entity.CreatedAt,
	}, nil
}

func (accountPersistence *AccountPersistence) Save(database *sql.DB, entity *Account) (*Account, error) {
	_, err := database.Exec(`
		INSERT INTO ledger__account(
			account__id,
			user__id,
			account__name,
			account_type__id,
			account__created_at)
		VALUES (?, ?, ?, ?, ?)`,
		entity.ID, entity.UserID, entity.Name, entity.Type, entity.CreatedAt)
	if err != nil {
		return nil, err
	}
	return entity, nil
}
