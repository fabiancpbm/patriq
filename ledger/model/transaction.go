package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID       uuid.UUID
	SourceID uuid.UUID
	TargetID uuid.UUID
	Date     time.Time
	Amount   float32
}
