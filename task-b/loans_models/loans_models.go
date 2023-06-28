package loans_models

import "time"

type Account struct {
	AccountID      int //Respresented by the MSIDN of the user
	Created_at     time.Time
	Unpaid_balance int
	LoanLimit      int
	Entries        []TransactionEntries
}

type TransactionEntries struct {
    ID int
    AmountAdded int
    Created_at     time.Time
    AccountID int

}