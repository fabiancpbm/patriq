package model

type AccountTypeID string

const (
	Asset     AccountTypeID = "asset"
	Liability AccountTypeID = "liability"
	Revenue   AccountTypeID = "revenue"
	Expense   AccountTypeID = "expense"
)

type AccountType struct {
	ID AccountTypeID
}
