package db

import (
	"github.com/go-openapi/errors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var sqlConnection *sql.DB

func ConnectDb() error {
	var err error
	if err != nil {
		return err
	}
	return nil
}

func AddUser(email, pin, name string) errors.Error {

}

func HasUser(email string) bool {

}
