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
	InsertThread(newThread models.NewThread, forum string) (LastID int, Err error)
	InsertUser(newUser models.NewUser, nickname string) (Err error)
	SelectUsersByNickname(nickname string) (user models.User, Err error)
	SelectUsersByNicknameOrEmail(email string, nickname string) (Users []models.User, Err error)
}
