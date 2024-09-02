package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	var database *sql.DB
	os.Setenv("LEDGER_DB_USER", "ledger_user")
	os.Setenv("LEDGER_DB_PASS", "12345678")
	config := mysql.Config{
		User:                 os.Getenv("LEDGER_DB_USER"),
		Passwd:               os.Getenv("LEDGER_DB_PASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "ledger",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	database, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}

	pingErr := database.Ping()
	if pingErr != nil {
		return nil, pingErr
	}
	logger := log.Default()
	logger.Print("Connected")
	return database, nil
}
