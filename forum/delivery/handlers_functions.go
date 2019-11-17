package delivery

import (
	"github.com/ValeryBMSTU/DB_TP/forum/usecase"
	"github.com/labstack/echo"
)

func (h *HandlersStruct) NewHandlers(e *echo.Echo, usecase usecase.UseInterface) error {
	h.PUsecase = usecase

	//e.GET("/", h.HandleEmpty)

	return nil
}
