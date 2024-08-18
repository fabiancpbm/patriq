package model

import (
	"github.com/google/uuid"
	"time"
)

type SourceType string

const (
	Statement SourceType = "statement"
	Invoice   SourceType = "invoice"
)

type Transaction struct {
	ID          uuid.UUID
	SourceID    string
	SourceType  SourceType
	Date        time.Time
	Amount      float32
	Description string
}
