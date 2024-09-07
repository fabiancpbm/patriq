package db

import "database/sql"

type JPA[Model any] interface {
	Save(database *sql.DB, model Model) (Model, error)
}
