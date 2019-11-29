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
	e.POST( "/api/forum/create", h.CreateForum)
	e.GET( "/api/forum/:slug/details", h.TakeForum)
	e.GET( "/api/forum/:slug/threads", h.TakeForumThreads)

	e.POST("/api/thread/:slug_or_id/create", h.CreatePosts)
	e.POST("/api/post/:id/details", h.ChangePost)
	e.GET( "/api/post/:id/details", h.TakePostByID)
	e.GET( "/api/thread/:slug_or_id/posts", h.TakePosts)

	e.POST( "/api/forum/:slug/create", h.CreateThread)
	e.POST( "/api/thread/:slug_or_id/details", h.ChangeThread)
	e.GET( "/api/thread/:slug_or_id/details", h.TakeThread)

	e.POST( "/api/user/:nickname/create", h.CreateUser)
	e.GET( "/api/forum/:slug/users", h.TakeUsersByForum)
	e.GET( "/api/user/:nickname/profile", h.TakeUser)
	e.POST("/api/user/:nickname/profile", h.ChangeUser)

	e.POST ("/api/thread/:slug_or_id/vote", h.CreateVote)

	e.GET( "/api/service/status", h.TakeStatus)

	e.POST( "/api/service/clear", h.Cleare)

	return nil
}