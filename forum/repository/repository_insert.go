package repository

import (
	"github.com/ValeryBMSTU/DB_TP/pkg/consts"
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
)

func (rep *ReposStruct) InsertForum(newForum models.NewForum) (Err error) {
	var lastID int
	err := rep.DataBase.QueryRow(consts.INSERTForum, newForum.Slug, newForum.Title,
		newForum.User).Scan(&lastID)

	if err != nil {
		return err
	}
	return nil
}

func (rep *ReposStruct) InsertThread(newThread models.NewThread, forum string) (LastID int, Err error) {
	var lastID int
	if newThread.Slug == "" {
		Err = rep.DataBase.QueryRow(consts.INSERTThread, newThread.Author, newThread.Created,
			newThread.Message, newThread.Title, forum).Scan(&lastID)
	} else {
		Err = rep.DataBase.QueryRow(consts.INSERTThreadWithSlug, newThread.Author, newThread.Created,
			newThread.Message, newThread.Title, forum, newThread.Slug).Scan(&lastID)
	}
	if Err != nil {
		return lastID, Err
	}
	return lastID, nil
}

func (rep *ReposStruct) InsertUser(newUser models.NewUser, nickname string) (Err error) {
	var lastID int
	err := rep.DataBase.QueryRow(consts.INSERTUser, newUser.About, newUser.Email,
		newUser.Fullname, nickname).Scan(&lastID)

	if err != nil {
		return err
	}
	return nil
}