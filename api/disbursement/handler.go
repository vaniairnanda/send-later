package disbursement

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
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

func (h *HTTPDisbursementHandler) CreateBatchDisbursement(context echo.Context) error {
	panic("implement me")
}
