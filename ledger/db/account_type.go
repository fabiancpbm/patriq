package db

import (
	"database/sql"

	"patriq.com.br/ledger/model"
)

type AccountType struct {
	ID model.AccountTypeID `db:"account_type__id"`
}

type AccountTypePersistence struct {
}

func (accountTypePersistence *AccountTypePersistence) ModelToEntity(model *model.AccountType) (*AccountType, error) {
	return &AccountType{
		ID: model.ID,
	}, nil
}

func (accountTypePersistence *AccountTypePersistence) EntityToModel(entity *AccountType) (*model.AccountType, error) {
	return &model.AccountType{
		ID: entity.ID,
	}, nil
}

func (accountTypePersistence *AccountTypePersistence) Save(database *sql.DB, entity *AccountType) (*AccountType, error) {
	_, err := database.Exec(`
		INSERT INTO ledger__account_type(
			account_type__id)
		VALUES (?)`,
		entity.ID)
	if err != nil {
		return nil, err
	}
	return entity, nil
}
