package model

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Name      string
	Type      AccountTypeID
	CreatedAt time.Time
}
