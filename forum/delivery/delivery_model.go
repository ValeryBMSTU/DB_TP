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
	e.GET( "/forum/:slug/threads", h.TakeForumThreads)

	e.POST("/thread/:slug_or_id/create", h.CreatePosts)
	e.POST("/post/:id/details", h.ChangePost)
	e.GET( "/thread/:slug_or_id/posts", h.TakePosts)

	e.POST( "/forum/:slug/create", h.CreateThread)
	e.POST( "/thread/:slug_or_id/details", h.ChangeThread)
	e.GET( "/thread/:slug_or_id/details", h.TakeThread)

	e.POST( "/user/:nickname/create", h.CreateUser)
	e.GET( "/forum/:slug/users", h.TakeUsersByForum)
	e.GET( "/user/:nickname/profile", h.TakeUser)
	e.POST("/user/:nickname/profile", h.ChangeUser)

	e.POST ("/thread/:slug_or_id/vote", h.CreateVote)


	return nil
}