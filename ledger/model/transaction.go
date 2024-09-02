package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID `db:"id"`
	SourceID  uuid.UUID `db:"source_id"`
	TargetID  uuid.UUID `db:"target_id"`
	Date      time.Time `db:"transaction_date"`
	Amount    float32   `db:"amount"`
	EventDate time.Time `db:"event_date"`
}
