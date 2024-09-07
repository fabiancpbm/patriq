package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AbstractAPI[DTO any, Model any] struct {
	Resource string
	Database *sql.DB
}

type API[DTO any, Model any] interface {
	DTOToModel(dto DTO) (*Model, error)
	Save(database *sql.DB, model Model) (*Model, error)
}


func (abstractApi *AbstractAPI[DTO, Model]) Post(c *gin.Context, api API[DTO, Model]) {
	var dto DTO
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, `{"error": "could not read json for [POST:`+
			abstractApi.Resource+
			`]"}`)
		return
	}

	model, err := api.DTOToModel(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, `{"error": "reading the body of `+
			abstractApi.Resource+
			`"}`)
		return
	}

	saved, err := create(abstractApi.Database, api, model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, `{"error": "`+
			err.Error()+
			`"}`)
		return
	}
	c.JSON(http.StatusCreated, saved)
}

func create[DTO any, Model any](database *sql.DB, objectConverter API[DTO, Model], model *Model) (*Model, error) {
	saved, err := objectConverter.Save(database, *model)
	if err != nil {
		return nil, err
	}

	return saved, nil
}
