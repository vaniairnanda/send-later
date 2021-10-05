package disbursement

import "time"

type Disbursement struct {
	ID                  uint64    `json:"id" gorm:"primary_key"`
	BatchDisbursementID uint64    `json:"batch_disbursement_id" gorm:"not null"`
	Amount              uint64    `json:"amount" gorm:"not null"`
	Description         string    `json:"description" gorm:"null"`
	BankCode            string    `json:"bank_code" gorm:"not null"`
	BankAccountName     string    `json:"bank_account_name" gorm:"not null"`
	BankAccountNumber   string    `json:"bank_account_number" gorm:"not null"`
	ExternalID          string    `json:"external_id" gorm:"not null"`
	CreatedAt           time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt           time.Time `json:"updated_at" gorm:"not null"`
}
