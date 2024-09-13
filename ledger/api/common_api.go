package api

import (
	"database/sql"

	"patriq.com.br/ledger/db"
	"patriq.com.br/ledger/logic"
)

type IApiPort[PostDtoIn any, Model any, PostDtoOut any] interface {
	PostDtoToModel(dto *PostDtoIn) (*Model, error)
	PostModelToDto(model *Model) (*PostDtoOut, error)
}

func Post[PostDtoIn any, Model any, Entity any, PostDtoOut any](
	iApi IApiPort[PostDtoIn, Model, PostDtoOut],
	iLogic logic.ILogic[Model],
	iPersistence db.IPersistencePort[Entity, Model],
	database *sql.DB,
	dto *PostDtoIn) (*PostDtoOut, error) {
	model, err := iApi.PostDtoToModel(dto)
	if err != nil {
		return nil, err
	}
	
	validModel, err := iLogic.Validate(model)
	if err != nil {
		return nil, err
	}

	savedModel, err := db.Save[Entity, Model](iPersistence, database, validModel)
	if err != nil {
		return nil, err
	}
	dtoOut, err := iApi.PostModelToDto(savedModel)
	if err != nil {
		return nil, err
	}
	return dtoOut, err
}
