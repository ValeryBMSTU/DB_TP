package usecase

import (
	"errors"
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
	"strconv"
)

func (use *UseStruct) GetForumsBySlug(slug string) (Forum []models.Forum, Err error) {
	forums, err := use.Rep.SelectForumsBySlug(slug)
	if err != nil {
		return forums, err
	}

	return forums,nil
}

func (use *UseStruct) GetPosts(slugOrID, limit, since, sort, desc string) (Posts *models.Posts, Err error) {

	var thread models.Thread
	id, err := strconv.Atoi(slugOrID)
	if err != nil {
		threads, err := use.Rep.SelectThreadsBySlug(slugOrID)
		if err != nil {
			return &models.Posts{}, err
		}
		if len(*threads) != 1 {
			return &models.Posts{}, errors.New("Can't find thread")
		}
		thread = *(*threads)[0]
	} else {
		threads, err := use.Rep.SelectThreadsByID(id)
		if err != nil {
			return &models.Posts{}, err
		}
		if len(*threads) != 1 {
			return &models.Posts{}, errors.New("Can't find thread")
		}
		thread = *(*threads)[0]
	}

	posts, err := use.Rep.SelectPosts(thread.ID, limit, since, sort, desc)
	if err != nil {
		return posts, err
	}

	return posts, err
}


func (use *UseStruct) GetThreadBySlug(slugOrID string) (Thread models.Thread, Err error) {
	var thread models.Thread
	id, err := strconv.Atoi(slugOrID)
	if err != nil {
		threads, err := use.Rep.SelectThreadsBySlug(slugOrID)
		if err != nil {
			return models.Thread{}, err
		}
		if len(*threads) != 1 {
			return models.Thread{}, errors.New("Can't find thread")
		}
		thread = *(*threads)[0]
	} else {
		threads, err := use.Rep.SelectThreadsByID(id)
		if err != nil {
			return models.Thread{}, err
		}
		if len(*threads) != 1 {
			return models.Thread{}, errors.New("Can't find thread")
		}
		thread = *(*threads)[0]
	}

	return thread,nil
}


func (use *UseStruct) GetThreadsByForum(forum string, limit string, since string, desc string) (Threads *models.Threads, Err error) {
	threads, err := use.Rep.SelectThreadsByForum(forum, limit, since, desc)
	if err != nil {
		return threads, err
	}

	return threads,nil
}

func (use *UseStruct) GetUsersByForum(slug, limit, desc string) (Users *models.Users, Err error) {
	users, err := use.Rep.SelectUsersByForum(slug, limit, desc)
	if err != nil {
		return users, err
	}

	return users,nil
}

func (use *UseStruct) GetUsersByNicknameOrEmail(email string, nickname string) (User []models.User, Err error) {
	users, err := use.Rep.SelectUsersByNicknameOrEmail(email, nickname)

	if err != nil {
		return users, err
	}

	return users,nil
}

func (use *UseStruct) GetUserByNickname(nickname string) (user models.User, Err error) {
	user, err := use.Rep.SelectUserByNickname(nickname)

	if err != nil {
		return user, err
	}

	return user,nil
}

func (use *UseStruct) GetUsersByEmail(email string) (User []models.User, Err error) {
	users, err := use.Rep.SelectUsersByEmail(email)

	if err != nil {
		return users, err
	}

	return users,nil
}