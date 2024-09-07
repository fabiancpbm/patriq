package model

// import (
// 	"github.com/google/uuid"
// 	"time"
// )

type AccountType string

const (
	Asset     AccountType = "asset"
	Liability AccountType = "liability"
	Revenue   AccountType = "revenue"
	Expense   AccountType = "expense"
)

// type Account struct {
// 	ID        uuid.UUID   `db: "account__id"`
// 	UserID    uuid.UUID   `db: "user__id"`
// 	Name      string      `db: "account__name"`
// 	Type      AccountType `db: "account__type"`
// 	CreatedAt time.Time   `db: "account__created_at"`
// }
