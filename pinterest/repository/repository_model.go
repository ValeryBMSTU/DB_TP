package repository

import (
	"database/sql"
)

type ReposStruct struct {
	connectionString string
	DataBase         *sql.DB
}

type ReposInterface interface {

}
