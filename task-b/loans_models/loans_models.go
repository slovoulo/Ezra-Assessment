package loans_models

import "time"

type Account struct {
	AccountID      int `gorm:"primaryKey"`//Respresented by the MSIDN of the user
	Created_at     time.Time  `gorm:"not null"`
	Updated_on 		time.Time
	Unpaid_balance int `gorm:"not null"`
	LoanLimit      int `gorm:"not null"`
	Entries        []TransactionEntries
}

type TransactionEntries struct {
    ID int `gorm:"primaryKey"`
    AmountAdded int `gorm:"not null"`
    Created_at     time.Time `gorm:"not null"`
    AccountID int `gorm:"not null"`
	TransactionType string `gorm:"not null"`

}