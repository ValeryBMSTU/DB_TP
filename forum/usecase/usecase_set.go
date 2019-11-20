package usecase

import (
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
	"github.com/lib/pq"
	"strconv"
)

func (use *UseStruct) SetThread(changeThread models.ChangeThread, slugOrID string) (Thread models.Thread, Err error) {
	thread := models.Thread{}
	id, err := strconv.Atoi(slugOrID)
	if err != nil {
		threads, err := use.Rep.SelectThreadsBySlug(slugOrID)
		if err != nil {
			return models.Thread{}, err
		}
		if len(*threads) != 1 {
			return models.Thread{}, nil
		}
		thread = *(*threads)[0]
		id = (*threads)[0].ID
	} else {
		threads, err := use.Rep.SelectThreadsByID(id)
		if err != nil {
			return models.Thread{}, err
		}
		if len(*threads) != 1 {
			return models.Thread{}, nil
		}
		thread = *(*threads)[0]
		id = (*threads)[0].ID
	}

	if changeThread.Message == "" {
		changeThread.Message = thread.Message
	} else {
		thread.Message = changeThread.Message
	}
	if changeThread.Title == "" {
		changeThread.Title = thread.Title
	} else {
		thread.Title = changeThread.Title
	}

	if err := use.Rep.UpdateThread(changeThread, id); err != nil {
		return models.Thread{}, err
	}



	return thread, nil
}

func (use *UseStruct) SetUser(newProfile models.NewUser, nickname string) (User models.User, Err error) {
	curentUser, err := use.Rep.SelectUserByNickname(nickname)
	if err != nil {
		return models.User{}, err
	}

	if newProfile.Email == "" {
		newProfile.Email = curentUser.Email
	}
	if newProfile.About == "" {
		newProfile.About = curentUser.About
	}
	if newProfile.Fullname == "" {
		newProfile.Fullname = curentUser.Fullname
	}

	if err := use.Rep.UpdateUser(newProfile, nickname); err != nil {
		return models.User{}, err
	}

	user := models.User{
		About:   newProfile.About,
		Email:    newProfile.Email,
		Fullname: newProfile.Fullname,
		Nickname: nickname,
	}

	return user,nil
}

func (use *UseStruct) SetVote(newVote models.NewVote, slugOrID string) (Thread models.Thread, Err error) {
	_, err := use.Rep.SelectUserByNickname(newVote.Nickname)
	if err != nil {
		return models.Thread{}, err
	}

	var thread models.Thread
	id, err := strconv.Atoi(slugOrID)
	if err != nil {
		threads, err := use.Rep.SelectThreadsBySlug(slugOrID)
		if err != nil {
			return models.Thread{}, err
		}
		if len(*threads) != 1 {
			return models.Thread{}, nil
		}
		id = (*threads)[0].ID
		thread = *(*threads)[0]
	} else {
		threads, err := use.Rep.SelectThreadsByID(id)
		if err != nil {
			return models.Thread{}, err
		}
		if len(*threads) != 1 {
			return models.Thread{}, nil
		}
		id = (*threads)[0].ID
		thread = *(*threads)[0]
	}

	err = use.Rep.InsertVote(newVote, id)
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if !ok {
			return models.Thread{}, err
		}
		if pqErr.Code == "23505" {
			err = use.Rep.UpdateVote(newVote, id)
		} else {
			return models.Thread{}, err
		}
	}
	if newVote.Voice == 1 {
		thread.Votes = thread.Votes + newVote.Voice
	} else {
		thread.Votes = thread.Votes - 2
	}

	return  thread, nil
}