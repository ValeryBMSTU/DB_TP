package delivery

import (
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

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