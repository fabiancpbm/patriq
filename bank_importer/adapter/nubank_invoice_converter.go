package adapter

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"patriq.com.br/bankimporter/model"
)

type NubankInvoiceConverter struct{}

func (*NubankInvoiceConverter) Convert(line []string) (*model.Transaction, error) {
	date, err := time.Parse(time.DateOnly, line[0])
	if err != nil {
		return nil, err
	}
	
	value, err := strconv.ParseFloat(line[3], 32)
	if err != nil {
		return nil, err
	}
	amount := float32(value)

	return &model.Transaction{
		ID:          uuid.New(),
		SourceID:    "",
		SourceType:  model.Invoice,
		Date:        date,
		Amount:      amount,
		Description: line[2],
	}, nil
}
