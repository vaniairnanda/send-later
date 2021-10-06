package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	mock_repository "github.com/vaniairnanda/send-later/mocks/api/disbursement"
	"github.com/vaniairnanda/send-later/model/disbursement"
	"github.com/vaniairnanda/send-later/model/dto"
	"github.com/vaniairnanda/send-later/model/enum"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type TestingHandlerSuite struct {
	suite.Suite
	*require.Assertions
	mock sqlmock.Sqlmock
	ctrl             *gomock.Controller
	Handler          *HTTPDisbursementHandler
	disbursementRepo *mock_repository.MockDisbursementRepository
	batchRepo        *mock_repository.MockBatchRepository
	dbDisbursement   *gorm.DB
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(TestingHandlerSuite))
}
func (s *TestingHandlerSuite) SetupTest() {

	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)


	dialect := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	s.dbDisbursement, err = gorm.Open(dialect, &gorm.Config{})
	require.NoError(s.T(), err)

	s.Assertions = require.New(s.T())

	s.ctrl = gomock.NewController(s.T())
	s.disbursementRepo = mock_repository.NewMockDisbursementRepository(s.ctrl)
	s.batchRepo = mock_repository.NewMockBatchRepository(s.ctrl)

	s.Handler = NewHTTPHandler(
		s.dbDisbursement,
		s.batchRepo,
		s.disbursementRepo,
	)

}

func (s *TestingHandlerSuite) TearDownTest() {
	s.ctrl.Finish()
}
func (s *TestingHandlerSuite) TestHTTPDisbursementHandler_CreateBatchDisbursement() {
	s.Run("when success", func() {
		ctx := context.Background()
		scheduledDateString := "2021-10-20"
		scheduledDate, _ := time.Parse("2006-01-02", scheduledDateString)

		request := dto.CreateDisbursement{
			Reference:     "",
			ScheduledDate: &scheduledDateString,
			ClientID:      1,
			CountryCode:   "Asia/Jakarta",
			IsSendLater:   true,
			Disbursements: []dto.Disbursement{
				{
					ExternalID:        "",
					Amount:            10000,
					BankCode:          "1",
					BankAccountName:   "test",
					BankAccountNumber: "123445",
					Description:       "test",
				},
			},
		}

		reqByte, _ := json.Marshal(request)

		ec := echo.New()

		httpReq := httptest.NewRequest(
			http.MethodPost,
			"/batch_disbursements",
			strings.NewReader(string(reqByte)),
		)
		httpReq.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)

		httpRes := httptest.NewRecorder()
		echoCtx := ec.NewContext(httpReq, httpRes)
		//db := gorm.DB{}
		//
		//
		daoBatch := disbursement.BatchDisbursement{
			ClientID:            request.ClientID,
			Reference:           &request.Reference,
			ScheduledDate:       &scheduledDate,
			CountryCode:         request.CountryCode,
			IsSendLater:         request.IsSendLater,
			TotalUploadedAmount: 10000,
			TotalUploadedCount:  1,
			Status:              enum.NEEDS_APPROVAL,
		}

		daoItems := []disbursement.Disbursement{
			{
				ExternalID:        "",
				Amount:            10000,
				BankCode:          "1",
				BankAccountName:   "test",
				BankAccountNumber: "123445",
				Description:       "test",
			},
		}

		s.mock.ExpectBegin()
		s.batchRepo.EXPECT().Store(ctx, s.dbDisbursement, daoBatch).Return(daoBatch, nil)
		s.disbursementRepo.EXPECT().BulkStore(ctx, s.dbDisbursement, daoItems).Return(int64(1), nil)
		err := s.Handler.CreateBatchDisbursement(echoCtx)

		s.Nil(err)
	})
}
