# patriq

## System Design

```mermaid
  C4Context
  title patriq Design System

  Person(user, "User", "The user that want to handle your financial life.")
  
  System_Ext(openFinance, "Open Finance API", "The open API from the Open Banking iniciative in Brasil")

  Enterprise_Boundary(b0, "patriq Software System") {

    System(mobileApp, "Mobile App", "The mobile App for the user experience.")

    System(gatewayApi, "Gateway API", "Responsible for render and interpret the content for the front-end.")

    System(ledgerApi, "Ledger API", "Handles with the User Accounts, Account Types, Financial Institutions, and Transactions.")

    System(userApi, "User API", "Handles with the User information and sessions.")

    System(bankImporterApi, "Bank Importer API", "API to get information from banking extract files.")

    SystemDb(ledgerDb, "Ledger DB", "A relational database to store  FinancialInstitution, Account, AccountType, and Transaction")

    SystemDb(userDb, "User DB", "A relational database to store Users.")

    SystemDb(userS3, "User S3", "An S3 for users Blob files.")

    Rel(userApi, userDb, "CRUD users")
    
    Rel(userApi, userS3, "CRUD users' pictures")

    Rel(bankImporterApi, ledgerApi, "post(transactions)")

    Rel(ledgerApi, ledgerDb, "CRUD ledger")

    Rel(user, mobileApp, "Register and login user, set accounts, bankings, and transactions, extracts, projects")

    BiRel(mobileApp, gatewayApi, "render product side information.")

    Rel(gatewayApi, ledgerApi, "post(accounts), post(financial-institutions), post(transactions), get(extracts), get(financial-institutions)")

    Rel(gatewayApi, userApi, "post(users), post(sessions), post(picture), get(users)")

    Rel(ledgerApi, openFinance, "Get transactions")
  }
```

<!-- SystemQueue(SystemF, "Banking System F Queue", "A system of the bank.")

    Component(C, "CC", "CCzinho")

    BiRel(openFinance, ledgerDb, "Uses")
    Rel(ledger, openFinance, "Sends e-mails", "SMTP") -->

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
### Import transactions

```bash
curl -v POST localhost:8080/transactions  -d '
{"sourceId":"d9fb7cd4-f650-4389-819b-b35429dfbfb2",
 "targetId":"262056a2-5f69-4903-8149-f108a8e3b6d9",
 "amount":10.5,
 "date":"2024-09-01T00:00:00"}'
```
