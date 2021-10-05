package handler

import (
	"context"
	"github.com/labstack/echo"
	"github.com/vaniairnanda/send-later/api/disbursement/batchRepository"
	"github.com/vaniairnanda/send-later/api/disbursement/disbursementRepository"
	"github.com/vaniairnanda/send-later/config"
	"github.com/vaniairnanda/send-later/model/constant"
	"github.com/vaniairnanda/send-later/model/dto"
	"github.com/vaniairnanda/send-later/shared"
	"gorm.io/gorm"
	"net/http"
)

type HTTPDisbursementHandler struct {
	DBDisbursement   *gorm.DB
}

func NewHTTPHandler(dbDisbursement *gorm.DB) *HTTPDisbursementHandler {
	return &HTTPDisbursementHandler{
		DBDisbursement: dbDisbursement,
	}
}

func (h *HTTPDisbursementHandler) Mount(group *echo.Group) {
	group.POST("/", h.CreateBatchDisbursement)
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
	// validate user role from business service
	if err := shared.MockBusinessServiceGetUserLogin(userLogin); err != nil {
		return err
	}
	storeBatch := payloadData.ToStoreBatch()
	tx := config.GetDBDisbursement().Begin()

	result, err := batchRepository.Store(ctx, tx, storeBatch)
	if err != nil {
		tx.Rollback()
		return err
	}

	storeItems := dto.ToStoreItems(payloadData.Disbursements, storeBatch.ID)
	rowsAffected, err := disbursementRepository.BulkStore(ctx, tx, storeItems)
	if err != nil || rowsAffected == 0 {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return c.JSON(http.StatusOK, result)
}


