package usecase

import "github.com/ValeryBMSTU/DB_TP/pkg/models"

func (use *UseStruct) GetUsersByNicknameOrEmail(email string, nickname string) (User []models.User, Err error) {
	users, err := use.Rep.SelectUsersByNicknameOrEmail(email, nickname)

	if err != nil {
		return users, err
	}

	return users,nil
}

func (use *UseStruct) GetUserByNickname(nickname string) (user models.User, Err error) {
	user, err := use.Rep.SelectUsersByNickname(nickname)

	if err != nil {
		return user, err
	}

	return user,nil
}