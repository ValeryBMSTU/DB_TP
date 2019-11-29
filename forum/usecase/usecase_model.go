package usecase

import (
	"github.com/ValeryBMSTU/DB_TP/forum/repository"
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
	"sync"
)


type UseStruct struct {
	Rep repository.ReposInterface
	Mu          *sync.Mutex
}

func (USC *UseStruct) NewUseCase(mu *sync.Mutex, rep repository.ReposInterface) error {
	USC.Mu = mu
	USC.Rep = rep
	return nil
}

type UseInterface interface {
	AddForum(newForum models.NewForum) (forum models.Forum, Err error)
	GetForumsBySlug(slug string) (forum []models.Forum, Err error)

	AddPosts(newPosts models.NewPosts, slug_or_id string) (posts models.Posts, Err error)
	SetPost(changePost models.ChangePost, postID int) (Post models.Post, Err error)
	GetPostByID(ID int, related string) (Post models.PostDetails, Err error)
	GetPosts(slugOrID, limit, since, sort, desc string) (posts *models.Posts, Err error)
	//GetPostDetails(postID int) (post models.Post, Err error)

	AddThread(newThread models.NewThread, forum string) (thread models.Thread, Err error)
	SetThread(changeThread models.ChangeThread, slugOrID string) (Thread models.Thread, Err error)
	GetThreadBySlug(slugOrID string) (Thread models.Thread, Err error)
	GetThreadsByForum(forum string, limit string, since string, desc string) (Threads *models.Threads, Err error)

	AddUser(newUser models.NewUser, nickname string) (user models.User, Err error)
	GetUsersByForum(slug, limit, since, desc string) (Users *models.Users, Err error)
	GetUserByNickname(nickname string) (user models.User, Err error)
	GetUsersByEmail(email string) (user []models.User, Err error)
	GetUsersByNicknameOrEmail(email string, nickname string) (user []models.User, Err error)
	SetUser(newProfile models.NewUser, nickname string) (user models.User, Err error)

	SetVote(newVote models.NewVote, slugOrID string) (Thread models.Thread, Err error)

	GetStatus() (Status models.Status, Errr error)

	Cleare() (Err error)
}
