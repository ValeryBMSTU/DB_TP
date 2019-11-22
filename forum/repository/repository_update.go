package repository

import (
	"github.com/ValeryBMSTU/DB_TP/pkg/consts"
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
)

func (rep *ReposStruct) UpdatePost(changePost models.ChangePost, postID int) (Err error) {
	rows, err := rep.DataBase.Query(consts.UPDATEPostByID, changePost.Message, true, postID)
	defer rows.Close()

	if err != nil {
		return err
	}
	return nil
}

func (rep *ReposStruct) UpdateThread(changeThread models.ChangeThread, id int) (Err error) {
	rows, err := rep.DataBase.Query(consts.UPDATEThreadByID, changeThread.Message,
		changeThread.Title, id)
	defer rows.Close()

	if err != nil {
		return err
	}
	return nil
}

func (rep *ReposStruct) UpdateUser(newProfile models.NewUser, nickname string) (Err error) {
	rows, err := rep.DataBase.Query(consts.UPDATEUserByNickname, newProfile.About, newProfile.Email,
		newProfile.Fullname, nickname)
	defer rows.Close()

	if err != nil {
		return err
	}
	return nil
}

func (rep *ReposStruct) UpdateVote(newVote models.NewVote, threadID int) (Err error) {
	rows, err := rep.DataBase.Query(consts.UPDATEVote, newVote.Voice, newVote.Nickname,
		threadID)
	defer rows.Close()

	if err != nil {
		return err
	}
	return nil
}
