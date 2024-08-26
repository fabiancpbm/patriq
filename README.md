# patriq

## bank-importer

### Import the bank transactions

```bash
curl -v POST localhost:8080/triggers  -d '
{"basePath":"/Users/fabian.brandao/Documents/patriq/bankimporter/input/",
 "year":2024,
 "month":8,
 "day":1,
 "bank":"nubank",
 "account":"154250440",
 "type":"statement"}'
```
- `type`: 
  - `statement`: the type used for debit extract
  - `invoice`: the type used for credit card extract

## ledger

```mermaid
classDiagram
  class AccountType {
    ASSET
    LIABILITY
    REVENUE
    EXPENSE
  }

  class Account {
    ID uuid.UUID
    Name string
    CreatedAt time.Time
  }

  class Transaction {
    ID uuid.UUID
    Date time.Time
    Amount float32
  }

  class FinancialInstitution {
    ID uuid.UUID
    AccountNumber string
  }

  class User {
    ID uuid.UUID
    Name string
    Email string
    CreatedAt time.Time
  }

  Account "0..*" -->  "1" AccountType: Type
  Account "0..*" --> "1" User: UserID
  Transaction "0..*" --> "1" Account: SourceID
  Transaction "0..*" --> "1" Account: TargetID
  FinancialInstitution "1" o--> "1..*" Account: Accounts
  FinancialInstitution "0..*" --> "1..*" User: UserID
```
