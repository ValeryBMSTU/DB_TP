package delivery

import (
	"encoding/json"
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
	"github.com/labstack/echo"
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

	forum, err := h.Use.AddForum(newForum)
	if err != nil {
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

	user, err := h.Use.AddUser(newUser, ctx.Param("nickname"))
	if err != nil {
		return err
	}

	if err := ctx.JSON(201, user); err != nil {
		return err
	}

	return nil
}


