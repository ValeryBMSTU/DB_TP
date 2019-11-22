package usecase

import (
	"errors"
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
	"strconv"
	"time"
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

func (use *UseStruct) AddPosts(newPosts models.NewPosts, slug_or_id string) (Posts models.Posts, Err error) {
	var forum string
	id, err := strconv.Atoi(slug_or_id)
	if err != nil {
		threads, err := use.Rep.SelectThreadsBySlug(slug_or_id)
		if err != nil {
			return models.Posts{}, err
		}
		if len(*threads) != 1 {
			return models.Posts{}, errors.New("Can't find thread")
		}
		forum = (*threads)[0].Forum
		id = (*threads)[0].ID
	} else {
		threads, err := use.Rep.SelectThreadsByID(id)
		if err != nil {
			return models.Posts{}, err
		}
		if len(*threads) != 1 {
			return models.Posts{},  errors.New("Can't find thread")
		}
		forum = (*threads)[0].Forum
		id = (*threads)[0].ID
	}

	posts := models.Posts{}
	created := time.Now()

	for _, newPost := range newPosts {
		if newPost.Parent != 0 {
			_, err := use.Rep.SelectPostByIDThreadID(newPost.Parent, id)
			if err != nil {
				return models.Posts{}, err
			}
		}

		lastID, threadID, err := use.Rep.InsertPost(*newPost, id, forum, created)
		if err != nil {
			return models.Posts{}, err
		}
		post := models.Post{
			Author:   newPost.Author,
			Created:  "",
			Forum:    forum,
			ID:       lastID,
			IsEdited: false,
			Message:  newPost.Message,
			Parent:   newPost.Parent,
			Thread:   threadID,
		}
		posts = append(posts, &post)
	}
	return  posts, nil
}

func (use *UseStruct) AddThread(newThread models.NewThread, forum string) (Thread models.Thread, Err error) {

	threads, err := use.Rep.SelectThreadsBySlug(newThread.Slug)
	if len(*threads) > 0 {
		return *(*threads)[0], errors.New("conflict")
	}
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

func (use *UseStruct) Cleare() (Err error) {
	if err := use.Rep.Cleare(); err != nil {
		return err
	}


	return nil
}