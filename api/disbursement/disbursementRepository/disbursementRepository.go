package disbursementRepository

import (
	"context"
	dao "github.com/vaniairnanda/send-later/model/disbursement"
	"gorm.io/gorm"
)

type Repository interface {
	BulkStore(ctx context.Context, db *gorm.DB, data []dao.Disbursement) (int64, error)
}


func BulkStore(ctx context.Context, db *gorm.DB,
	data []dao.Disbursement) (int64, error) {

	result := db.CreateInBatches(data, 10000)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
