package repository

import (
	"github.com/ValeryBMSTU/DB_TP/pkg/consts"
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
	"time"
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

func (rep *ReposStruct) InsertPost(newPost models.NewPost, id int, forum string, created time.Time) (LastID int, ThreadID int, Err error) {
	var lastID, threadID int
	err := rep.DataBase.QueryRow(consts.INSERTPost, newPost.Author, newPost.Message,
		newPost.Parent, id, forum).Scan(&lastID, &threadID)

	if err != nil {
		return lastID, threadID, err
	}
	return lastID, threadID,nil
}

func (rep *ReposStruct) InsertThread(newThread models.NewThread, forum string) (LastID int, Err error) {
	var lastID int
	if newThread.Slug == "" {
		if newThread.Created == "" {
			Err = rep.DataBase.QueryRow(consts.INSERTThreadWithoutCreated, newThread.Author,
				newThread.Message, newThread.Title, forum).Scan(&lastID)
		} else {
			Err = rep.DataBase.QueryRow(consts.INSERTThread, newThread.Author, newThread.Created,
				newThread.Message, newThread.Title, forum).Scan(&lastID)
		}
	} else {
		if newThread.Created == "" {
			Err = rep.DataBase.QueryRow(consts.INSERTThreadWithSlugWithoutCreated, newThread.Author,
				newThread.Message, newThread.Title, forum, newThread.Slug).Scan(&lastID)
		} else {
			Err = rep.DataBase.QueryRow(consts.INSERTThreadWithSlug, newThread.Author, newThread.Created,
				newThread.Message, newThread.Title, forum, newThread.Slug).Scan(&lastID)
		}
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

func (rep *ReposStruct) InsertVote(newVote models.NewVote, threadID int) (Err error) {
	var lastID int
	err := rep.DataBase.QueryRow(consts.INSERTVote, newVote.Nickname, newVote.Voice,
		threadID).Scan(&lastID)

	if err != nil {
		return err
	}
	return nil
}

func (rep *ReposStruct) Cleare() (Err error) {
	rows, err := rep.DataBase.Query(consts.CLEARE)
	defer rows.Close()
	if err != nil {
		return err
	}

	return nil
}