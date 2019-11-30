package repository

import (
	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
	"io/ioutil"
)
//
//func (RS *ReposStruct) DataBaseInit() error {
//	RS.connectionString = consts.ConnStr
//	var err error
//
//	RS.DataBase, err = sql.Open("postgres", consts.ConnStr)
//	if err != nil {
//		return err
//	}
//	RS.DataBase.SetMaxOpenConns(100)
//	err = RS.DataBase.Ping()
//	if err != nil {
//		return err
//	}
//
//	if err := RS.LoadSchemaSQL(); err != nil {
//		err, ok := err.(*pq.Error)
//		if !ok {
//			return err
//		}
//		if err.Code != pq.ErrorCode("42P06") {
//			return err
//		}
//	}
//
//	//RS.Cleare()
//
//	return nil
//}
//
//func (RS *ReposStruct) LoadSchemaSQL() (Err error) {
//	dbSchema := "sunrise_db.sql"
//
//	content, err := ioutil.ReadFile(dbSchema)
//	if err != nil {
//		return err
//	}
//	tx, err := RS.DataBase.Begin()
//	if err != nil {
//		return err
//	}
//	defer func() {
//		if err := tx.Rollback(); err != nil {
//			Err = errors.Wrap(Err, err.Error())
//		}
//	}()
//
//	if _, err = tx.Exec(string(content)); err != nil {
//		return err
//	}
//	if err := tx.Commit(); err != nil {
//		return err
//	}
//	return nil
//}
//
//func (RS *ReposStruct) NewDataBaseWorker() error {
//	RS.connectionString = consts.ConnStr
//	var err error
//
//	RS.DataBase, err = sql.Open("postgres", consts.ConnStr)
//	if err != nil {
//		return err
//	}
//	RS.DataBase.SetMaxOpenConns(100)
//	err = RS.DataBase.Ping()
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (RS *ReposStruct) CloseDB() error {
//	if err := RS.DataBase.Close(); err != nil {
//		return err
//	}
//	return nil
//}


//var connPool *pgx.ConnPool = nil

const maxConn = 2000
const dbSchema = "sunrise_db.sql"

func (RS *ReposStruct) DataBaseInit(psqURI string) error {
	if RS.DataBase != nil {
		return nil
	}
	config, err := pgx.ParseURI(psqURI)
	if err != nil {
		return err
	}

	RS.DataBase, err = pgx.NewConnPool(
		pgx.ConnPoolConfig{
			ConnConfig:     config,
			MaxConnections: maxConn,
		})
	if err != nil {
		return err
	}

	err = RS.LoadSchemaSQL()
	if err != nil {
		return err
	}

	return nil
}

func (RS *ReposStruct) LoadSchemaSQL() error {
	if RS.DataBase == nil {
		return pgx.ErrDeadConn
	}

	content, err := ioutil.ReadFile(dbSchema)
	if err != nil {
		return err
	}

	tx, err := RS.DataBase.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err = tx.Exec(string(content)); err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (RS *ReposStruct) Disconn() {
	if RS.DataBase != nil {
		RS.DataBase.Close()
		RS.DataBase = nil
	}
}

func (RS *ReposStruct) Query(sql string, args ...interface{}) (*pgx.Rows, error) {
	if RS.DataBase == nil {
		return nil, pgx.ErrDeadConn
	}
	return RS.DataBase.Query(sql, args...)
}

func (RS *ReposStruct) QueryRow(sql string, args ...interface{}) *pgx.Row {
	if RS.DataBase == nil {
		return nil
	}
	return RS.DataBase.QueryRow(sql, args...)
}

func (RS *ReposStruct) Exec(sql string, args ...interface{}) (pgx.CommandTag, error) {
	if RS.DataBase == nil {
		return "", pgx.ErrDeadConn
	}

	tx, err := RS.DataBase.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	tag, err := tx.Exec(sql, args...)
	if err != nil {
		return "", err
	}
	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return tag, nil
}