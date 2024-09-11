package db

import "database/sql"

type IPersistencePort[Entity any, Model any] interface {
	ModelToEntity(model *Model) (*Entity, error)
	EntityToModel(entity *Entity) (*Model, error)
	Save(database *sql.DB, entity *Entity) (*Entity, error)
}

func Save[Entity any, Model any](iPersistence IPersistencePort[Entity, Model], database *sql.DB, model *Model) (*Model, error) {
	entity, err := iPersistence.ModelToEntity(model)
	if err != nil {
		return nil, err
	}
	savedEntity, err := iPersistence.Save(database, entity)
	if err != nil {
		return nil, err
	}
	savedModel, err := iPersistence.EntityToModel(savedEntity)
	if err != nil {
		return nil, err
	}
	return savedModel, nil
}
