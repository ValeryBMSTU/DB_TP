package main

import (
	"github.com/ValeryBMSTU/DB_TP/forum/delivery"
	"github.com/ValeryBMSTU/DB_TP/forum/repository"
	use "github.com/ValeryBMSTU/DB_TP/forum/usecase"
	"github.com/ValeryBMSTU/DB_TP/pkg/consts"
	customMiddlewares "github.com/ValeryBMSTU/DB_TP/pkg/middlewares"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"sync"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: consts.LoggerFormat}))
	e.Use(customMiddlewares.PanicMiddleware)
	e.HTTPErrorHandler = customMiddlewares.CustomHTTPErrorHandler


	handlers := delivery.HandlersStruct{}
	var mutex sync.Mutex
	rep := repository.ReposStruct{}
	err := rep.DataBaseInit("postgresql://forum:forum@localhost:5432/forum")
	if err != nil {
		e.Logger.Errorf("repository error: %s", err)
	}

	useCase := use.UseStruct{}
	err = useCase.NewUseCase(&mutex, &rep)
	if err != nil {
		e.Logger.Errorf("server error: %s", err)
	}
	err = handlers.NewHandlers(e, &useCase)
	if err != nil {
		e.Logger.Errorf("server error: %s", err)
	}

	e.Logger.Warnf("start listening on %s", consts.HostAddress)
	if err := e.Start(consts.HostAddress); err != nil {
		e.Logger.Errorf("server error: %s", err)
	}

	e.Logger.Warnf("shutdown")
}
