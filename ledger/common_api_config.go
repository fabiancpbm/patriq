package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"patriq.com.br/ledger/api"
	"patriq.com.br/ledger/db"
	"patriq.com.br/ledger/logic"
)

type AbstractApiConfig[PostDtoIn any, Model any, Entity any, PostDtoOut any] struct {
	Resource    string
	Database    *sql.DB
	Api api.IApiPort[PostDtoIn, Model, PostDtoOut]
	Logic logic.ILogic[Model]
	Persistence db.IPersistencePort[Entity, Model]
}

func (apiConfig *AbstractApiConfig[PostDtoIn, Model, Entity, PostDtoOut]) Post(c *gin.Context) {
	var dto PostDtoIn
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, `{"error": "could not read json for [POST:`+
			apiConfig.Resource+
			`]"}`)
		return
	}

	saved, err := api.Post(apiConfig.Api, apiConfig.Logic, apiConfig.Persistence, apiConfig.Database, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, `{"error": "`+
			err.Error()+
			`"}`)
		return
	}
	c.JSON(http.StatusCreated, saved)
}
