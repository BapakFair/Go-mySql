package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username string = "root"
	password string = "123123123"
	database string = "db_movie"
)

var dsn = fmt.Sprintf("%v:%v@/%v", username, password, database)

func MySql() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}
	return db, nil
}
