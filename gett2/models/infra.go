package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
)

const (
	DB_CONNECTION_STRING = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
	DRIVER_NAME = "postgres"
)

func openConnection() *sql.DB {
	DB_HOST, _ := beego.GetConfig("String", "DB_HOST", "localhost")
	DB_PORT, _ := beego.GetConfig("String", "DB_PORT", "5432")
	DB_USER, _ := beego.GetConfig("String", "DB_USER", "postgres")
	DB_PASSWORD, _ := beego.GetConfig("String", "DB_PASSWORD", "admin")
	DB_NAME, _ := beego.GetConfig("String", "DB_NAME", "postgres")

	psqlInfo := fmt.Sprintf(DB_CONNECTION_STRING, DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open(DRIVER_NAME, psqlInfo)

	err = db.Ping()
	checkErr(err)

	return db
}

func closeConnection(db *sql.DB) {
	defer db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}