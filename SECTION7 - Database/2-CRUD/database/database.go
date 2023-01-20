package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Mysql driver
)

// Connect with database
func Connect() (*sql.DB, error) {
	stringKeyToConnect := "golang:golang@/devbook_golang_course?charset=utf8&parseTime=True&loc=Local"

	//Opens connection with DB
	database, err := sql.Open("mysql", stringKeyToConnect)
	if err != nil {
		return nil, err
	}

	//Check if the connection with DB is OK
	if err := database.Ping(); err != nil {
		return nil, err
	}

	//Return database and error(nil) if it's all OK
	return database, nil
}
