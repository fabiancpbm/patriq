package adapter

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"patriq.com.br/bankimporter/model"
)

type TransactionImporterImpl struct {
	Trigger model.Trigger
}

// importTransactions implements port.TransactionImporter.
func (i TransactionImporterImpl) ImportTransactions() ([]model.Transaction, error) {
	filePath := getFilePath(i.Trigger)
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var transactions []model.Transaction
	transactionConverter := initTransactionConverter(i.Trigger)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.New("Error while reading the file " + filePath)
		}
		if transactions == nil {
			transactions = make([]model.Transaction, 0)
			continue
		}

		transaction, err := transactionConverter.Convert(record)
		if err != nil {
			return nil, errors.New("Error while converting the line " + strings.Join(record, ", "))
		}
		transactions = append(transactions, *transaction)
	}
	return transactions, nil
}

func getFilePath(trigger model.Trigger) string {
	year := fmt.Sprintf("%02d", trigger.Year)
	month := fmt.Sprintf("%02d", trigger.Month)
	day := fmt.Sprintf("%02d", trigger.Day)
	triggerType := string(trigger.Type)

	fileName := strings.Join([]string{year, month, day, trigger.Bank, trigger.Account, triggerType}, "_") + ".csv"

	filePath := path.Join(trigger.BasePath, year, month, day, trigger.Bank, trigger.Account, triggerType, fileName)
	return filePath
}

func initTransactionConverter(trigger model.Trigger) TransactionConverter {
	switch context := trigger.Bank + "_" + string(trigger.Type); context {
	case "nubank_invoice":
		return &NubankInvoiceConverter{}
	case "nubank_statement":
		return &NubankStatementConverter{}
	default:
		return nil
	}
}
