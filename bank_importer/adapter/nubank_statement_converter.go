package adapter

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"patriq.com.br/bankimporter/model"
)

type NubankStatementConverter struct {
}

func (*NubankStatementConverter) Convert(line []string) (*model.Transaction, error) {
	date, err := time.Parse("02/01/2006", line[0])
	if err != nil {
		return nil, err
	}

	value, err := strconv.ParseFloat(line[1], 32)
	if err != nil {
		return nil, err
	}
	amount := float32(value)

	return &model.Transaction{
		ID:          uuid.New(),
		SourceID:    line[2],
		SourceType:  model.Statement,
		Date:        date,
		Amount:      amount,
		Description: line[3],
	}, nil
}
