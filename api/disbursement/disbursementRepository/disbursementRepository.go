package disbursementRepository

import (
	"context"
	"github.com/vaniairnanda/send-later/api/disbursement"
	dao "github.com/vaniairnanda/send-later/model/disbursement"
	"gorm.io/gorm"
)
type disbursementRepository struct {}


func NewRepository() disbursement.DisbursementRepository{
	return &disbursementRepository{}
}
func (r *disbursementRepository) BulkStore(ctx context.Context, db *gorm.DB,
	data []dao.Disbursement) (int64, error) {

	result := db.CreateInBatches(data, 10000)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
