package app

import (
	"fmt"
	"github.com/labstack/echo"
	echoMid "github.com/labstack/echo/middleware"
	"os"
	"strconv"
)

const DefaultPort = 8080

func (app *App) HTTPServerMain() {
	e := echo.New()
	e.Use(echoMid.Recover())
	e.Use(echoMid.CORS())

	e.Debug = true
	handlerGroup := e.Group("/batch-disbursements")
	app.DisbursementHandler.Mount(handlerGroup)


	// set REST port
	var port uint16
	if portEnv, ok := os.LookupEnv("SERVER_PORT"); ok {
		portInt, err := strconv.Atoi(portEnv)
		if err != nil {
			port = DefaultPort
		} else {
			port = uint16(portInt)
		}
	} else {
		port = DefaultPort
	}

	listenerPort := fmt.Sprintf(":%d", port)
	e.Logger.Fatal(e.Start(listenerPort))
}


