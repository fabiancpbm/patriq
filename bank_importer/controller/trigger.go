package controller

import (
	"fmt"
	"net/http"

	"patriq.com.br/bankimporter/model"
	"patriq.com.br/bankimporter/port"
)

func ProcessTrigger(importer port.TransactionImporter, httpClient *http.Client, trigger model.Trigger) error {
	transactions, err := importer.ImportTransactions()
	if err != nil {
		return err
	}

	for _, transaction := range transactions {
		// httpClient.Post()
		fmt.Println(transaction)
	}

	return nil
}
