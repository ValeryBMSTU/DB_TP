package usecase

import "github.com/ValeryBMSTU/DB_TP/pkg/models"

func (use *UseStruct) SetUser(newProfile models.NewUser, nickname string) (User models.User, Err error) {
	if err := use.Rep.UpdateUser(newProfile, nickname); err != nil {
		return models.User{}, err
	}

	user := models.User{
		About:   newProfile.About,
		Email:    newProfile.Email,
		Fullname: newProfile.Fullname,
		Nickname: nickname,
	}

	return user,nil
}
