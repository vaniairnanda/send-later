package handler

import (
	"github.com/labstack/echo"
	dao "github.com/vaniairnanda/send-later/model/disbursement"
	"github.com/vaniairnanda/send-later/model/dto"
	"gorm.io/gorm"
	"testing"
)

func TestHTTPDisbursementHandler_CreateBatchDisbursement(t *testing.T) {
	type fields struct {
		DBDisbursement *gorm.DB
	}
	type args struct {
		c       echo.Context
		payload dto.CreateDisbursement
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   dao.BatchDisbursement
		want1  error
	}{
		{
			name: "Success Case - Create Batch Disbursement",
			args: args{},

		}
	},
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HTTPDisbursementHandler{
				DBDisbursement: tt.fields.DBDisbursement,
			}
			if err := h.CreateBatchDisbursement(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CreateBatchDisbursement() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
