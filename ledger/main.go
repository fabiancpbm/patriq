package main

import (
	"github.com/gin-gonic/gin"
	"patriq.com.br/ledger/api"
	"patriq.com.br/ledger/db"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		panic(err)
	}
	dependences := &api.Dependences {
		Database: database,
	}
		
	router := gin.Default()
	router.POST("/transactions", dependences.PostTransaction)
	router.Run(":8080")
}
