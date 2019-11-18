package repository

import (
	"github.com/ValeryBMSTU/DB_TP/pkg/consts"
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
)

func (rep *ReposStruct) UpdateUser(newProfile models.NewUser, nickname string) (Err error) {
	rows, err := rep.DataBase.Query(consts.UPDATEUserByNickname, newProfile.About, newProfile.Email,
		newProfile.Fullname, nickname)
	defer rows.Close()

	if err != nil {
		return err
	}
	return nil
}
