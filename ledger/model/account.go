package model

import (
	"github.com/google/uuid"
	"time"
)

type AccountType string

const (
	Asset     AccountType = "asset"
	Liability AccountType = "liability"
	Revenue   AccountType = "revenue"
	Expense   AccountType = "expense"
)

type Account struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Name      string
	Type      AccountType
	CreatedAt time.Time
}
