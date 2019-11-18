package usecase

import "github.com/ValeryBMSTU/DB_TP/pkg/models"

func (use *UseStruct) SetUser(newProfile models.NewUser, nickname string) (User models.User, Err error) {
	curentUser, err := use.Rep.SelectUserByNickname(nickname)
	if err != nil {
		return models.User{}, err
	}

	if newProfile.Email == "" {
		newProfile.Email = curentUser.Email
	}
	if newProfile.About == "" {
		newProfile.About = curentUser.About
	}
	if newProfile.Fullname == "" {
		newProfile.Fullname = curentUser.Fullname
	}

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
