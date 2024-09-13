package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"patriq.com.br/ledger/api"
	"patriq.com.br/ledger/db"
	"patriq.com.br/ledger/logic"
	"patriq.com.br/ledger/model"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.POST(
		"/account-types",
		func(c *gin.Context) {
			getAccountTypeApiConfig(database).Post(c)
		})
	router.POST(
		"/accounts",
		func(c *gin.Context) {
			getAccountApiConfig(database).Post(c)
		})
	router.POST(
		"/transactions",
		func(c *gin.Context) {
			getTransactionApiConfig(database).Post(c)
		})

	router.Run(":8080")
}

func getAccountTypeApiConfig(database *sql.DB) *AbstractApiConfig[api.AccoutTypeIn, model.AccountType, db.AccountType, api.AccountTypeOut] {
	accountTypeApi := &api.AccountTypeAPI{}
	accountTypeLogic := &logic.AccountTypeLogic{}
	accountTypePersistence := &db.AccountTypePersistence{}
	return &AbstractApiConfig[api.AccoutTypeIn, model.AccountType, db.AccountType, api.AccountTypeOut]{
		Resource:    "account_type",
		Database:    database,
		Api:         accountTypeApi,
		Logic:       accountTypeLogic,
		Persistence: accountTypePersistence,
	}
}

func getAccountApiConfig(database *sql.DB) *AbstractApiConfig[api.AccountIn, model.Account, db.Account, api.AccountOut] {
	accountApi := &api.AccountAPI{}
	accountLogic := &logic.AccountLogic{}
	accountPersistence := &db.AccountPersistence{}
	return &AbstractApiConfig[api.AccountIn, model.Account, db.Account, api.AccountOut]{
		Resource:    "account",
		Database:    database,
		Api:         accountApi,
		Logic:       accountLogic,
		Persistence: accountPersistence,
	}
}

func getTransactionApiConfig(database *sql.DB) *AbstractApiConfig[api.TransactionIn, model.Transaction, db.Transaction, api.TransactionOut] {
	transactionApi := &api.TransactionAPI{}
	transactionLogic := &logic.TransactionLogic{}
	transactionPersistence := &db.TransactionPersistence{}
	return &AbstractApiConfig[api.TransactionIn, model.Transaction, db.Transaction, api.TransactionOut]{
		Resource:    "transaction",
		Database:    database,
		Api:         transactionApi,
		Logic:       transactionLogic,
		Persistence: transactionPersistence,
	}
}
