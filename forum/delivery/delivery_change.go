package delivery

import (
	"encoding/json"
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"strconv"
)

func (h *HandlersStruct) ChangePost(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()

	ctx.Response().Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(ctx.Request().Body)

	changePost:= models.ChangePost{}

	if err := decoder.Decode(&changePost); err != nil {
		return err
	}

	ID := ctx.Param( "id")
	postID, err := strconv.Atoi(ID)
	if err != nil {
		return err
	}

	post, err := h.Use.SetPost(changePost, postID)
	if err != nil {
		if err := ctx.JSON(404, models.Error{"Can't find post"}); err != nil {
			return err
		}
		return nil
	}

	if err := ctx.JSON(200, post); err != nil {
		return err
	}

	return nil
}

func (h *HandlersStruct) ChangeThread(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()

	ctx.Response().Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(ctx.Request().Body)

	changeThread := models.ChangeThread{}

	if err := decoder.Decode(&changeThread); err != nil {
		return err
	}

	slugOrID := ctx.Param( "slug_or_id")

	thread, err := h.Use.SetThread(changeThread, slugOrID)
	if err != nil {
		if err := ctx.JSON(404, models.Error{"Can't find thread"}); err != nil {
			return err
		}
		return nil
	}

	if err := ctx.JSON(200, thread); err != nil {
		return err
	}

	return nil
}

func (h *HandlersStruct) ChangeUser(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()

	ctx.Response().Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(ctx.Request().Body)

	newProfile:= models.NewUser{}

	if err := decoder.Decode(&newProfile); err != nil {
		return err
	}


	users, err := h.Use.GetUsersByEmail(newProfile.Email)
	if err != nil {
		return err
	}
	if len(users) > 0 && !(len(users) == 1 &&
		users[0].Nickname == ctx.Param("nickname")){
		if err := ctx.JSON(409, models.Error{"Conflict"}); err != nil {
			return err
		}

		return nil
	}

	user, err := h.Use.SetUser(newProfile, ctx.Param("nickname"))
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