package delivery

import (
	"encoding/json"
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
	"github.com/labstack/echo"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

func (h *HandlersStruct) CreateForum(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()

	ctx.Response().Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(ctx.Request().Body)

	newForum := models.NewForum{}

	if err := decoder.Decode(&newForum); err != nil {
		return err
	}

	forums, err := h.Use.GetForumsBySlug(newForum.Slug)
	if err != nil {
		return err
	}
	if len(forums) > 0 {
		if err := ctx.JSON(409, forums[0]); err != nil {
			return err
		}
		return nil
	}

	forum, err := h.Use.AddForum(newForum)
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if !ok {
			return err
		}
		if pqErr.Code == "23503" {
			if err := ctx.JSON(404, models.Error{"Can't find user"}); err != nil {
				return err
			}
			return nil
		}
		return err
	}

	if err := ctx.JSON(201, forum); err != nil {
		return err
	}

	return nil
}

func (h *HandlersStruct) CreateThread(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()

	ctx.Response().Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(ctx.Request().Body)

	newThread:= models.NewThread{}

	if err := decoder.Decode(&newThread); err != nil {
		return err
	}

	thread, err := h.Use.AddThread(newThread, ctx.Param("slug"))
	if err != nil {
		return err
	}

	if err := ctx.JSON(201, thread); err != nil {
		return err
	}

	return nil
}

func (h *HandlersStruct) CreateUser(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()

	ctx.Response().Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(ctx.Request().Body)

	newUser:= models.NewUser{}

	if err := decoder.Decode(&newUser); err != nil {
		return err
	}

	users, err := h.Use.GetUsersByNicknameOrEmail(newUser.Email,ctx.Param("nickname"))
	if err != nil {
		return err
	}
	if len(users) > 0 {
		if err := ctx.JSON(409, users); err != nil {
			return err
		}
		return nil
	}

	user, err := h.Use.AddUser(newUser, ctx.Param("nickname"))
	if err != nil {
		return err
	}


	//if err != nil {
	//	pqErr, ok := err.(*pq.Error)
	//	if !ok {
	//		return err
	//	}
	//	if pqErr.Message == `повторяющееся значение ключа нарушает ограничение уникальности "user_nickname_uindex"` ||
	//		pqErr.Code == "23505" {
	//		users, err := h.Use.GetUsersByNicknameOrEmail(newUser.Email,ctx.Param("nickname"))
	//		if err != nil {
	//			return err
	//		}
	//		if err := ctx.JSON(409, users); err != nil {
	//			return err
	//		}
	//		return nil
	//	}
	//	return err
	//}

	if err := ctx.JSON(201, user); err != nil {
		return err
	}

	return nil
}


