package delivery

import (
	"encoding/json"
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

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


	users, err := h.Use.GetUsersByNicknameOrEmail(newProfile.Email, "")
	if err != nil {
		return err
	}
	if len(users) > 0 {
		if err := ctx.JSON(409, models.Error{"Conflict"}); err != nil {
			return err
		}

		return nil
	}

	user, err := h.Use.SetUser(newProfile, ctx.Param("nickname"))
	if err != nil {
		return err
	}

	if err := ctx.JSON(200, user); err != nil {
		return err
	}

	return nil
}