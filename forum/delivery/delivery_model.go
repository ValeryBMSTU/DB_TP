package delivery

import (
	"github.com/ValeryBMSTU/DB_TP/forum/usecase"
	"github.com/labstack/echo"
)

type HandlersStruct struct {
	Use usecase.UseInterface
}

func (h *HandlersStruct) NewHandlers(e *echo.Echo, usecase usecase.UseInterface) error {
	h.Use = usecase

	//e.GET("/", h.HandleEmpty)
	e.POST( "/forum/create", h.CreateForum)
	e.GET( "/forum/:slug/details", h.TakeForum)

	e.POST( "/forum/:slug/create", h.CreateThread)

	e.POST( "/user/:nickname/create", h.CreateUser)
	e.GET( "/user/:nickname/profile", h.TakeUser)
	e.POST("user/:nickname/profile", h.ChangeUser)


	return nil
}