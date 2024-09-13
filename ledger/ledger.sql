DROP DATABASE IF EXISTS ledger;
CREATE DATABASE ledger;
USE ledger;

-- ========================================================
-- ACCOUNT
-- ========================================================

CREATE TABLE IF NOT EXISTS ledger__account_type (
	account_type__id VARCHAR(255) PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS ledger__account (
	account__id CHAR(36) PRIMARY KEY,
	user__id CHAR(36) NOT NULL,
	account__name VARCHAR(255) NOT NULL,
	account_type__id VARCHAR(255) NOT NULL,
	account__created_at DATETIME,
	FOREIGN KEY (account_type__id) REFERENCES ledger__account_type(account_type__id)
);

-- ========================================================
-- FINANCIAL INSTITUTION
-- ========================================================

CREATE TABLE IF NOT EXISTS ledger__finnancial_institution (
	financial_institution__code VARCHAR(3) PRIMARY KEY,
	financial_institution__name VARCHAR(255) NOT NULL
);

-- ========================================================
-- FINANCIAL PRODUCT
-- ========================================================

CREATE TABLE IF NOT EXISTS ledger__finnancial_product_type (
	financial_product_type__id VARCHAR(255) PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS ledger__financial_product (
	financial_product__id CHAR(36) PRIMARY KEY,
	financial_product_type__id VARCHAR(255) NOT NULL,
	financial_institution__code VARCHAR(3) NOT NULL,
	user__id CHAR(36) NOT NULL,
	account__id CHAR(36) NOT NULL,
	FOREIGN KEY (financial_product_type__id) REFERENCES ledger__finnancial_product_type(financial_product_type__id),
	FOREIGN KEY (account__id) REFERENCES ledger__account(account__id)
);

-- ========================================================
-- TRANSACTION
-- ========================================================

CREATE TABLE IF NOT EXISTS ledger__category (
	category__id CHAR(36) PRIMARY KEY,
	category__name VARCHAR(255) NOT NULL,
	user__id CHAR(36) NOT NULL
);

CREATE TABLE IF NOT EXISTS ledger__transaction (
    transaction__id CHAR(36) PRIMARY KEY,
    transaction__source_id CHAR(36) NOT NULL,
	transaction__target_id CHAR(36) NOT NULL,
	transaction__date DATETIME NOT NULL,
	transaction__event_date DATETIME NOT NULL,
	transaction__amount FLOAT NOT NULL CHECK (transaction__amount >= 0),
	category__id CHAR(36),
	FOREIGN KEY (transaction__source_id) REFERENCES ledger__account(account__id),
	FOREIGN KEY (transaction__target_id) REFERENCES ledger__account(account__id),
	FOREIGN KEY (category__id) REFERENCES ledger__category(category__id)
);
