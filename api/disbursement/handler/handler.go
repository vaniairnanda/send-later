package handler

import (
	"context"
	"errors"
	"github.com/labstack/echo"
	"github.com/vaniairnanda/send-later/api/disbursement"
	"github.com/vaniairnanda/send-later/api/disbursement/batchRepository"
	"github.com/vaniairnanda/send-later/api/disbursement/disbursementRepository"
	"github.com/vaniairnanda/send-later/config"
	"github.com/vaniairnanda/send-later/model/constant"
	"github.com/vaniairnanda/send-later/model/dto"
	"github.com/vaniairnanda/send-later/model/enum"
	"github.com/vaniairnanda/send-later/shared"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type HTTPDisbursementHandler struct {
	DBDisbursement *gorm.DB
	BatchRepo disbursement.BatchRepository
	DisbursementRepo disbursement.DisbursementRepository
}

func NewHTTPHandler(dbDisbursement *gorm.DB) *HTTPDisbursementHandler {
	batchRepo := batchRepository.NewRepository()
	disbursementRepo := disbursementRepository.NewRepository()
	return &HTTPDisbursementHandler{
		DBDisbursement: dbDisbursement,
		BatchRepo: batchRepo,
		DisbursementRepo: disbursementRepo,
	}
}

func (h *HTTPDisbursementHandler) Mount(group *echo.Group) {
	group.POST("", h.CreateBatchDisbursement)
	group.PATCH("/approve/:batchId", h.ApproveBatchDisbursement)
}

func (h *HTTPDisbursementHandler) CreateBatchDisbursement(c echo.Context) error {
	// placeholder userlogin
	c.Set(constant.USER_LOGIN_KEY, "uploader")
	payloadData := new(dto.CreateDisbursement)
	ctx := context.Background()

	if err := c.Bind(payloadData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := shared.ValidateCreate(*payloadData); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			err.Error(),
		)
	}

	userLogin := shared.GetUserLogin(c)
	// mock validate user role from business service
	if err := shared.MockBusinessServiceGetUserLogin(userLogin); err != nil {
		return err
	}
	storeBatch := payloadData.ToStoreBatch()
	tx := config.GetDBDisbursement().Begin()

	result, err := h.BatchRepo.Store(ctx, tx, storeBatch)
	if err != nil {
		tx.Rollback()
		return err
	}

	storeItems := dto.ToStoreItems(payloadData.Disbursements, storeBatch.ID)
	rowsAffected, err := h.DisbursementRepo.BulkStore(ctx, tx, storeItems)
	if err != nil || rowsAffected == 0 {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return c.JSON(http.StatusOK, result)
}

func (h *HTTPDisbursementHandler) ApproveBatchDisbursement(c echo.Context) error {
	ctx := context.Background()
	// placeholder userlogin
	c.Set(constant.USER_LOGIN_KEY, "approver")

	batchID, err := strconv.Atoi(c.Param("batchId"))
	if err != nil {
		return errors.New("invalid value for batch id")
	}
	payloadData := new(dto.ApproveDisbursement)
	if err := c.Bind(payloadData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := shared.ValidateApprove(*payloadData); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			err.Error(),
		)
	}

	var updateData = map[string]interface{}{}

	if payloadData.IsInstantDisbursement {
		updateData = map[string]interface{}{
			"status":         enum.PROCESSING,
		}

	} else {
		scheduledDate, _ := time.Parse("2006-01-02", *payloadData.NewScheduledDate)
		updateData = map[string]interface{}{
			"status":         enum.APPROVED,
			"scheduled_date": scheduledDate,
		}
	}

	db := config.GetDBDisbursement()
	result, err := h.BatchRepo.PatchByID(ctx, db, updateData, uint64(batchID))
	if err != nil{
		return err
	}

	if payloadData.IsInstantDisbursement {
		go shared.PublishEventDisbursementApply(result)
	}

	return c.JSON(http.StatusOK, result)

}
