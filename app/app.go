package app

import (
	"github.com/vaniairnanda/send-later/api/disbursement"
	"gorm.io/gorm"
	"sync"
)

type App struct {
	DisbursementHandler *disbursement.HTTPDisbursementHandler
}

func MakeHandler(dbDisbursement *gorm.DB) *App {
	disbursementNewHandler := disbursement.NewHTTPHandler(dbDisbursement)

	return &App{
		DisbursementHandler: disbursementNewHandler,
	}
}

func (a *App) Start() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		a.HTTPServerMain()
	}()

	wg.Wait()
}
