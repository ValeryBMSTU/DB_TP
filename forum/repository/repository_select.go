package repository

import (
	"errors"
	"github.com/ValeryBMSTU/DB_TP/pkg/consts"
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
)

func (rep *ReposStruct) SelectForumsBySlug(slug string) (Forum []models.Forum, Err error) {
	var forums []models.Forum
	rows, err := rep.DataBase.Query(consts.SELECTForumsBySlug, slug)
	defer rows.Close()
	if err != nil {
		return forums, err
	}

	scanForum := models.Forum{}
	for rows.Next() {
		err := rows.Scan(&scanForum.Posts, &scanForum.Slug, &scanForum.Thread,
			&scanForum.Title, &scanForum.User)
		if err != nil {
			return forums, err
		}
		forums = append(forums, scanForum)
	}

	return forums, nil
}

func (rep *ReposStruct) SelectUsersByNicknameOrEmail(email string, nickname string) (Users []models.User, Err error) {
	var users []models.User
	rows, err := rep.DataBase.Query(consts.SELECTUsersByNicknameOrEmail, email, nickname)
	defer rows.Close()
	if err != nil {
		return users, err
	}

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

func (rep *ReposStruct) SelectUserByNickname(nickname string) (user models.User, Err error) {
	var users []models.User
	rows, err := rep.DataBase.Query(consts.SELECTUsersByNickname, nickname)
	defer rows.Close()
	if err != nil {
		return models.User{}, err
	}

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

func (rep *ReposStruct) SelectUsersByEmail(email string) (Users []models.User, Err error) {
	var users []models.User
	rows, err := rep.DataBase.Query(consts.SELECTUsersByEmail, email)
	defer rows.Close()
	if err != nil {
		return users, err
	}

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

