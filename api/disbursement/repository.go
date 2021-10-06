package disbursement

import (
	"context"
	dao "github.com/vaniairnanda/send-later/model/disbursement"
	"gorm.io/gorm"
)

type BatchRepository interface {
	Store(ctx context.Context, db *gorm.DB, data dao.BatchDisbursement) (dao.BatchDisbursement, error)
	PatchByID(ctx context.Context, db *gorm.DB,
		data map[string]interface{}, id uint64) (*dao.BatchDisbursement, error)
}

type DisbursementRepository interface {
	BulkStore(ctx context.Context, db *gorm.DB, data []dao.Disbursement) (int64, error)
}
