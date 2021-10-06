package app

import (
	"github.com/vaniairnanda/send-later/api/disbursement/batchRepository"
	"github.com/vaniairnanda/send-later/api/disbursement/disbursementRepository"
	"github.com/vaniairnanda/send-later/api/disbursement/handler"
	"gorm.io/gorm"
	"sync"
)

type App struct {
	DisbursementHandler *handler.HTTPDisbursementHandler
}

func MakeHandler(dbDisbursement *gorm.DB) *App {
	batchRepo := batchRepository.NewRepository()
	disbursementRepo := disbursementRepository.NewRepository()
	disbursementNewHandler := handler.NewHTTPHandler(dbDisbursement, batchRepo, disbursementRepo)

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
