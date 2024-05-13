package dblayer

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DBconfig() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(mysqlDatabase:3306)/urlShortnerDB?parseTime=true")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
