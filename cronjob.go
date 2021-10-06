package main

import (
	"fmt"
	"github.com/vaniairnanda/send-later/config"
	"github.com/vaniairnanda/send-later/config/kafka"
	"github.com/vaniairnanda/send-later/environment"
	"github.com/vaniairnanda/send-later/model/disbursement"
	"github.com/vaniairnanda/send-later/model/enum"
	"github.com/vaniairnanda/send-later/shared"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
	"time"
)
type Interface interface {
	JobApprovalExpired()
	JobScheduledBatchDisbursement()

	InitializePublisher()
}

type job struct {
	DB *gorm.DB
}

func NewJob() Interface {
	return &job{}
}


func (job *job) InitializePublisher() {
	// connect to kafka
	kafkaProducer, err := kafka.Configure(strings.Split(kafkaBrokerUrl, ","), kafkaClientId, kafkaTopic)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to configure kafka: %v\n", err)
		return
	}
	defer kafkaProducer.Close()
}

func (job *job) JobApprovalExpired() {
	db := config.GetDBDisbursement()
	var batchResult []disbursement.BatchDisbursement

	err := db.Model(&disbursement.BatchDisbursement{}).
		Where("is_send_later = ?", true).
		Where("status = ?", enum.NEEDS_APPROVAL).
		Where("scheduled_date < current_date AT time zone country_code").
		Find(&batchResult).Error

	if err != nil {
		zap.S().Errorf("Error query get batch disbursement expired. Error= %v", err.Error())
		return
	}

	updateData := map[string]interface{}{
		"status":         enum.EXPIRED,
		"updated_at":      time.Now(),
	}

	for _, item := range batchResult {
		if err = db.Model(&disbursement.BatchDisbursement{}).
			Where("id = ?", item.ID).
			Updates(updateData).Error; err != nil {
			zap.S().Errorf("Error query update expired batch disbursement. Error= %v", err.Error())
			continue
		}

		go shared.PublishEventApprovalExpired(item) // publish message to be consumed by NotificationService
	}
}


func (job *job) JobScheduledBatchDisbursement() {
	db := config.GetDBDisbursement()
	var batchResult []disbursement.BatchDisbursement
	env := environment.Load()
	timeNow := time.Now()
	scheduledDisbursementTime := strings.Split(env.ScheduledDisbursementTime, ":")
	scheduledDisbursementHour, _ := strconv.Atoi(scheduledDisbursementTime[0])
	scheduledDisbursementMinutes, _ := strconv.Atoi(scheduledDisbursementTime[1])
	defaultDate := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), scheduledDisbursementHour, scheduledDisbursementMinutes,0, 0, time.UTC)

	err := db.Model(&disbursement.BatchDisbursement{}).
		Where("is_send_later = ?", true).
		Where("status = ?", enum.APPROVED).
		Where("scheduled_date = current_date AT time zone country_code").
		Where("now() AT time zone country_code >= ?", defaultDate).
		Find(&batchResult).Error

	if err != nil {
		zap.S().Errorf("Error query get batch disbursement eligible. Error= %v", err.Error())
		return
	}

	updateData := map[string]interface{}{
		"status":         enum.PROCESSING,
		"updated_at":      time.Now(),
	}

	for _, item := range batchResult {
		if err = db.Model(&disbursement.BatchDisbursement{}).
			Where("id = ?", item.ID).
			Updates(updateData).Error; err != nil {
			zap.S().Errorf("Error query update eligible batch disbursement. Error= %v", err.Error())
			continue
		}

		go shared.PublishEventDisbursementApply(item) // publish message to be consumed by itself
	}
}

