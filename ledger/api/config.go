package api

import "database/sql"

type Dependences struct {
	Database *sql.DB
}
