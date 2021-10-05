package shared

import (
	"errors"
	"github.com/labstack/echo"
	"github.com/vaniairnanda/send-later/model/constant"
	"github.com/vaniairnanda/send-later/model/dto"
	"time"
)

func ValidateCreate(payload dto.CreateDisbursement) error {

	if payload.IsSendLater && payload.ScheduledDate == nil {
		return errors.New("scheduled date is required for disbursement type Send Later")
	}

	if payload.ClientID == 0 {
		return errors.New("client id is required")

	}

	scheduledDate, err := time.Parse("2006-01-02", *payload.ScheduledDate)
	if err != nil {
		return errors.New("invalid date format")
	}

	if scheduledDate.Before(time.Now()) {
		return errors.New("scheduled date must be greater than today")
	}

	return nil
}

func GetUserLogin(c echo.Context) string {
	val := c.Get(constant.USER_LOGIN_KEY)
	login := val.(string)
	return login
}