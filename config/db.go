package config

import (
	"database/sql"
)

func Connect() *sql.DB {
	dbDriver := "mysql"
	// dbUserName := "sql6521053"
	// dbPass := "Y16yPKWige"
	// dbHost := "sql6.freesqldatabase.com"
	// dbPort := "3306"

	db, err := sql.Open(dbDriver, "sql6521053:Y16yPKWige@tcp(sql6.freesqldatabase.com:3306)/sql6521053")
	if err != nil {
		panic(err.Error())
	}
	return db
}
