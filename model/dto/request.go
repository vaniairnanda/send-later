package dto

import (
	"github.com/vaniairnanda/send-later/model/disbursement"
	"github.com/vaniairnanda/send-later/model/enum"
	"time"
)

type (
	CreateDisbursement struct {
		Reference     string         `json:"reference"`
		ScheduledDate *string        `json:"scheduled_date"`
		ClientID      uint64         `json:"client_id" validate:"required"`
		CountryCode   string         `json:"country_code" validate:"required"`
		IsSendLater   bool           `json:"is_send_later"`
		Disbursements []Disbursement `json:"disbursements" validate:"required"`
	}

	Disbursement struct {
		ExternalID        string `json:"external_id"`
		Amount            uint64 `json:"amount"`
		BankCode          string `json:"bank_code"`
		BankAccountName   string `json:"bank_account_name"`
		BankAccountNumber string `json:"bank_account_number"`
		Description       string `json:"description" `
	}
)

func (data CreateDisbursement) ToStoreBatch() disbursement.BatchDisbursement {
	scheduledDate, _ := time.Parse("2006-01-02", *data.ScheduledDate)
	totalAmount := calculateTotalAmount(data.Disbursements)
	return disbursement.BatchDisbursement{
		Reference:           &data.Reference,
		ScheduledDate:       &scheduledDate,
		CountryCode:         data.CountryCode,
		IsSendLater:         data.IsSendLater,
		TotalUploadedAmount: totalAmount,
		TotalUploadedCount:  uint64(len(data.Disbursements)),
		Status:              enum.NEEDS_APPROVAL,
		ClientID:            data.ClientID,
	}
}

func ToStoreItems(disbursements []Disbursement, batchID uint64) (result []disbursement.Disbursement) {
	for _, item := range disbursements {
		disbursementItem := disbursement.Disbursement{
			BatchDisbursementID: batchID,
			Amount:              item.Amount,
			Description:         item.Description,
			BankCode:            item.BankCode,
			BankAccountName:     item.BankAccountName,
			BankAccountNumber:   item.BankAccountNumber,
			ExternalID:          item.ExternalID,
		}

		result = append(result, disbursementItem)
	}

	return result
}

func calculateTotalAmount(disbursements []Disbursement) (total uint64) {
	for _, item := range disbursements {
		total += item.Amount
	}

	return total
}
