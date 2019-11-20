package delivery

import (
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func (h *HandlersStruct) TakeForum(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()

	ctx.Response().Header().Set("Content-Type", "application/json")



	forums, err := h.Use.GetForumsBySlug(ctx.Param("slug"))

	if err != nil || len(forums) == 0 {
		if err := ctx.JSON(404, models.Error{"Can't find forum by slug"}); err != nil {
			return err
		}
		return nil
	}

	if err := ctx.JSON(200, forums[0]); err != nil {
		return err
	}

	return nil
}

func (h *HandlersStruct) TakeForumThreads(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()

	ctx.Response().Header().Set("Content-Type", "application/json")

	limit := ctx.QueryParam("limit")
	since := ctx.QueryParam("since")
	desc := ctx.QueryParam("desc")

	if limit == "" {
		limit = "100"
	}
	if desc == "" {
		desc = "false"
	}

	forums, err := h.Use.GetForumsBySlug(ctx.Param("slug"))

	if err != nil || len(forums) == 0 {
		if err := ctx.JSON(404, models.Error{"Can't find forum by slug"}); err != nil {
			return err
		}
		return nil
	}

	threads, err := h.Use.GetThreadsByForum(ctx.Param("slug"), limit, since, desc)
	if err != nil {
		return err
	}

	if err := ctx.JSON(200, threads); err != nil {
		return err
	}

	return nil
}

func (h *HandlersStruct) TakeThread(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()

	ctx.Response().Header().Set("Content-Type", "application/json")


	slug_or_id := ctx.Param("slug_or_id")

	thread, err := h.Use.GetThreadBySlug(slug_or_id)

	if err != nil {
		if err := ctx.JSON(404, models.Error{"Can't thread"}); err != nil {
			return err
		}
		return nil
	}

	if err := ctx.JSON(200, thread); err != nil {
		return err
	}

	return nil
}

func (h *HandlersStruct) TakeUser(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()

	ctx.Response().Header().Set("Content-Type", "application/json")



	user, err := h.Use.GetUserByNickname(ctx.Param("nickname"))

	if err != nil {
		if err.Error() == "Can't find user by nickname" {
			if err := ctx.JSON(404, models.Error{err.Error()}); err != nil {
				return err
			}
			return nil
		}
		return err
	}

	if err := ctx.JSON(200, user); err != nil {
		return err
	}

	return nil
}

func (h *HandlersStruct) TakeUsersByForum(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()

	ctx.Response().Header().Set("Content-Type", "application/json")

	slug := ctx.Param("slug")

	forums, err := h.Use.GetForumsBySlug(slug)
	if len(forums) != 1 || err != nil{
		if err := ctx.JSON(404, models.Error{"Can't find forum by slug"}); err != nil {
			return err
		}
	}

	limit := ctx.QueryParam("limit")
	desc := ctx.QueryParam( "desc")

	if limit == "" {
		limit = "100"
	}
	if desc == "" {
		desc = "false"
	}

	users, err := h.Use.GetUsersByForum(slug, limit, desc)
	if err != nil {
		return err
	}

	if err := ctx.JSON(200, users); err != nil {
		return err
	}

	return nil
}

func (h *HandlersStruct) TakePosts(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()

	ctx.Response().Header().Set("Content-Type", "application/json")

	slugOrID := ctx.Param("slug_or_id")

	limit := ctx.QueryParam("limit")
	since := ctx.QueryParam("since")
	sort := ctx.QueryParam("sort")
	desc := ctx.QueryParam( "desc")

	if limit == "" {
		limit = "100"
	}

	if sort == "" {
		sort = "flat"
	}
	if desc == "" {
		desc = "false"
	}
	if since == "" {
		if desc == "false" {
			since = "0"
		} else {
			since = "999999999"
		}
	}

	posts, err := h.Use.GetPosts(slugOrID, limit, since, sort, desc)
	if err != nil {
		return err
	}

	if err := ctx.JSON(200, posts); err != nil {
		return err
	}

	return nil
}
