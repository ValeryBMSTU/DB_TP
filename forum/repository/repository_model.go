package repository

import (
	"database/sql"
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
)

type ReposStruct struct {
	connectionString string
	DataBase         *sql.DB
}

type ReposInterface interface {
	InsertForum(newForum models.NewForum) (Err error)
	SelectForumsBySlug(slug string) (forum []models.Forum, Err error)

	InsertThread(newThread models.NewThread, forum string) (LastID int, Err error)
	SelectThreadsByForum(forum string, limit string, since string, desc string) (Threads *models.Threads, Err error)

	InsertUser(newUser models.NewUser, nickname string) (Err error)
	SelectUserByNickname(nickname string) (user models.User, Err error)
	SelectUsersByEmail(email string) (Users []models.User, Err error)
	SelectUsersByNicknameOrEmail(email string, nickname string) (Users []models.User, Err error)
	UpdateUser(newProfile models.NewUser, nickname string) (Err error)
}
