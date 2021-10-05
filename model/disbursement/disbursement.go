package disbursement

type Disbursement struct {
	ID                  uint64 `json:"id" gorm:"primary_key"`
	BatchDisbursementID uint64 `json:"batch_disbursement_id" gorm:"not null"`
	Value               uint64 `json:"value" gorm:"not null"`
	BankCode            string `json:"bank_code" gorm:"not null"`
	BankAccountName     string `json:"bank_account_name" gorm:"not null"`
	BankNumber          string `json:"bank_number" gorm:"not null"`
	ExternalID          string `json:"external_id" gorm:"not null"`
}
