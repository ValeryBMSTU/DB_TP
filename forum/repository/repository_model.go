package repository

import (
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
	"github.com/jackc/pgx"
	"time"
)

type ReposStruct struct {
	connectionString string
	DataBase         *pgx.ConnPool
}

type ReposInterface interface {
	InsertForum(newForum models.NewForum) (Err error)
	SelectForumsBySlug(slug string) (forum []models.Forum, Err error)

	InsertPost(newPost models.NewPost, id int, forum string, created time.Time) (LastID int, Thread int, Err error)
	SelectPostByID(ID int) (Post models.Post, Err error)
	SelectPostByIDThreadID(ID int, threadID int) (Post models.Post, Err error)
	SelectPosts(threadID int, limit, since, sort, desc string) (Posts *models.Posts, Err error)
	//SelectPostByID(postID int) (Post models.Post, Err error)
	UpdatePost(changePost models.ChangePost, postID int) (Err error)

	InsertThread(newThread models.NewThread, forum string) (LastID int, Err error)
	UpdateThread(changeThread models.ChangeThread, id int) (Err error)
	SelectThreadsBySlug(slug string) (threads *models.Threads, Err error)
	SelectThreadsByID(id int) (threads *models.Threads, Err error)
	SelectThreadsByForum(forum string, limit string, since string, desc string) (Threads *models.Threads, Err error)

	InsertUser(newUser models.NewUser, nickname string) (Err error)
	SelectUsersByForum(slug, limit, since, desc string) (Users *models.Users, Err error)
	SelectUserByNickname(nickname string) (user models.User, Err error)
	SelectUsersByEmail(email string) (Users []models.User, Err error)
	SelectUsersByNicknameOrEmail(email string, nickname string) (Users []models.User, Err error)
	UpdateUser(newProfile models.NewUser, nickname string) (Err error)

	InsertVote(newVote models.NewVote, id int) (Err error)
	UpdateVote(newVote models.NewVote, id int) (Err error)

	SelectStatus() (Status models.Status, Err error)

	Cleare() (Err error)
}
