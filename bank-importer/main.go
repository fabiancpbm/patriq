package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"patriq.com.br/bankimporter/adapter"
	"patriq.com.br/bankimporter/controller"
	"patriq.com.br/bankimporter/model"
)

func main() {
	router := gin.Default()
	router.POST("/triggers", postTriggers)

	log.SetPrefix("[patriq.com.br/bankimporter]: ")
	log.SetFlags(0)

	router.Run("localhost:8080")
}

func postTriggers(c *gin.Context) {
	var trigger model.Trigger

	if err := c.BindJSON(&trigger); err != nil {
		c.JSON(http.StatusBadRequest, `{"error": "could not read json for [POST:triggers]"}`)
		return
	}

	httpClient := &http.Client{}

	transactionImporter := adapter.TransactionImporterImpl{Trigger: trigger}
	err := controller.ProcessTrigger(transactionImporter, httpClient, trigger)
	if err != nil {
		c.JSON(http.StatusInternalServerError, `{"error": "` + err.Error() + `"}`)
		return
	}
	c.JSON(http.StatusProcessing, trigger)
}
