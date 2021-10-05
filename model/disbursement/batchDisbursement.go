package disbursement

import (
	"github.com/vaniairnanda/send-later/model/enum"
	"time"
)

type BatchDisbursement struct {
	ID                  uint64      `json:"id" gorm:"primary_key;auto_increment:true;"`
	ClientID            uint64      `json:"client_id" gorm:"not null"`
	Reference           *string     `json:"reference" gorm:"null"`
	ScheduledDate       *time.Time  `json:"scheduled_date" gorm:"null"`
	CountryCode         string      `json:"country_code" gorm:"not null"`
	IsSendLater         bool        `json:"is_send_later" gorm:"not null"`
	ApprovedAt          *time.Time  `json:"approved_at" gorm:"null"`
	TotalUploadedAmount uint64      `json:"total_uploaded_amount" gorm:"not null"`
	TotalUploadedCount  uint64      `json:"total_uploaded_count" gorm:"not null"`
	Status              enum.Status `json:"status" gorm:"not null"`
	CreatedAt           time.Time   `json:"created_at" gorm:"not null"`
	UpdatedAt           time.Time   `json:"updated_at" gorm:"not null"`
}
