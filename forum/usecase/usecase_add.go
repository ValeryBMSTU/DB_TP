package usecase

import (
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
)

func (use *UseStruct) AddForum(newForum models.NewForum) (Forum models.Forum, Err error) {
	user, err := use.Rep.SelectUserByNickname(newForum.User)
	if err != nil {
		return models.Forum{}, err
	}

	newForum.User = user.Nickname

	if err := use.Rep.InsertForum(newForum); err != nil {
		return models.Forum{}, err
	}

	forum := models.Forum{
		Posts:  0,
		Slug:   newForum.Slug,
		Thread: 0,
		Title:  newForum.Title,
		User:   newForum.User,
	}

	return forum ,nil
}

func (use *UseStruct) AddThread(newThread models.NewThread, forum string) (Thread models.Thread, Err error) {
	lastID, err := use.Rep.InsertThread(newThread, forum)
	if err != nil {
		return models.Thread{}, err
	}

	thread := models.Thread{
		Author:  newThread.Author,
		Created: newThread.Created,
		Forum:   forum,
		ID:      lastID,
		Message: newThread.Message,
		Slug:    newThread.Slug,
		Title:   newThread.Title,
		Votes:   0,
	}

	return thread, nil
}

func (use *UseStruct) AddUser(newUser models.NewUser, nickname string) (User models.User, Err error) {
	if err := use.Rep.InsertUser(newUser, nickname); err != nil {
		return models.User{}, err
	}

	user := models.User{
		About:   newUser.About,
		Email:    newUser.Email,
		Fullname: newUser.Fullname,
		Nickname: nickname,
	}

	return user,nil
}