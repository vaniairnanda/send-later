package batchRepository

import (
	"context"
	dao "github.com/vaniairnanda/send-later/model/disbursement"
	"gorm.io/gorm"
)

type Repository interface {
	Store(ctx context.Context, db *gorm.DB, data dao.BatchDisbursement) (dao.BatchDisbursement, error)
}

func Store(ctx context.Context, db *gorm.DB,
	data dao.BatchDisbursement) (dao.BatchDisbursement, error) {

	result := db.Create(&data)
	if result.Error != nil {
		return dao.BatchDisbursement{}, result.Error
	}

	return data, nil
}

func PatchByID(ctx context.Context, db *gorm.DB,
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
