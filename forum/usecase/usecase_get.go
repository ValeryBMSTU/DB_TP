package usecase

import (
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
)

func (use *UseStruct) GetForumsBySlug(slug string) (Forum []models.Forum, Err error) {
	forums, err := use.Rep.SelectForumsBySlug(slug)
	if err != nil {
		return forums, err
	}

	return forums,nil
}

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

func (use *UseStruct) GetUsersByEmail(email string) (User []models.User, Err error) {
	users, err := use.Rep.SelectUsersByEmail(email)

	if err != nil {
		return users, err
	}

	return users,nil
}