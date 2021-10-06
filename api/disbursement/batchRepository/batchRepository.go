package batchRepository

import (
	"context"
	"github.com/vaniairnanda/send-later/api/disbursement"
	dao "github.com/vaniairnanda/send-later/model/disbursement"
	"gorm.io/gorm"
)


type batchRepository struct {}


func NewRepository() disbursement.BatchRepository {
	return &batchRepository{}
}

func (r *batchRepository) Store(ctx context.Context, db *gorm.DB,
	data dao.BatchDisbursement) (dao.BatchDisbursement, error) {

	result := db.Create(&data)
	if result.Error != nil {
		return dao.BatchDisbursement{}, result.Error
	}

	return data, nil
}

func (r *batchRepository) PatchByID(ctx context.Context, db *gorm.DB,
	data map[string]interface{}, id uint64) (*dao.BatchDisbursement, error) {
	var result dao.BatchDisbursement
	if err := db.
		First(&result, id).
		Model(&result).
		Updates(data).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
