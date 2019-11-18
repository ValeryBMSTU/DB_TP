package repository

import (
	"errors"
	"github.com/ValeryBMSTU/DB_TP/pkg/consts"
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
)

func (rep *ReposStruct) SelectUsersByNicknameOrEmail(email string, nickname string) (Users []models.User, Err error) {
	var users []models.User
	rows, err := rep.DataBase.Query(consts.SELECTUsersByNicknameOrEmail, email, nickname)
	if err != nil {
		return users, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			Err = err
		}
	}()

	scanUser := models.User{}
	for rows.Next() {
		err := rows.Scan(&scanUser.About, &scanUser.Email, &scanUser.Fullname,
			&scanUser.Nickname)
		if err != nil {
			return users, err
		}
		users = append(users, scanUser)
	}
	return users, nil
}

func (rep *ReposStruct) SelectUsersByNickname(nickname string) (user models.User, Err error) {
	var users []models.User
	rows, err := rep.DataBase.Query(consts.SELECTUsersByNickname, nickname)
	if err != nil {
		return models.User{}, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			Err = err
		}
	}()

	scanUser := models.User{}
	for rows.Next() {
		err := rows.Scan(&scanUser.About, &scanUser.Email, &scanUser.Fullname,
			&scanUser.Nickname)
		if err != nil {
			return models.User{}, err
		}
		users = append(users, scanUser)
	}

	if len(users) == 0 {
		return models.User{}, errors.New("Can't find user by nickname")
	}
	return users[0], nil
}

