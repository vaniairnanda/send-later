package environment

import (
	"github.com/joho/godotenv"
	"os"
)

type Env struct {
	DBUsername                 string
	DBPassword                 string
	DBHost                     string
	DBPort                     string
	DBName                     string
	MarkApprovalExpired        string
	ScheduledBatchDisbursement string
	ScheduledDisbursementTime  string
	SendApprovalReminder       string
}

func Load() *Env {
	godotenv.Load(".env")

	return &Env{
		DBUsername:                 os.Getenv("DB_USERNAME"),
		DBPassword:                 os.Getenv("DB_PASSWORD"),
		DBHost:                     os.Getenv("DB_HOST"),
		DBPort:                     os.Getenv("DB_PORT"),
		DBName:                     os.Getenv("DB_NAME"),
		MarkApprovalExpired:        os.Getenv("MARK_APPROVAL_EXPIRED"),
		ScheduledBatchDisbursement: os.Getenv("SCHEDULED_BATCH_DISBURSEMENT"),
		ScheduledDisbursementTime:  os.Getenv("SCHEDULED_DISBURSEMENT_TIME"),
		SendApprovalReminder:       os.Getenv("SEND_APPROVAL_REMINDER"),
	}
}
