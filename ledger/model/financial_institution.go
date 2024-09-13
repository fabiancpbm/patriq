package model

import "github.com/google/uuid"

type FinancialInstitution struct {
	ID uuid.UUID
	AccountNumber string
	Accounts []uuid.UUID
}
