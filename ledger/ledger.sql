DROP DATABASE IF EXISTS ledger;
CREATE DATABASE ledger;
USE ledger;

CREATE TABLE transaction (
    id CHAR(36) PRIMARY KEY,
    source_id CHAR(36) NOT NULL,
	target_id CHAR(36) NOT NULL,
	transaction_date DATETIME,
	event_date DATETIME,
	amount FLOAT
)
