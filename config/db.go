package config

import (
	"database/sql"
)

func Connect() *sql.DB {
	dbDriver := "mysql"
	// dbUserName := "sql6519232"
	// dbPass := "bsLgAUakKa"
	// dbHost := "sql6.freesqldatabase.com"
	// dbPort := "3306"

	db, err := sql.Open(dbDriver, "sql6519232:bsLgAUakKa@tcp(sql6.freesqldatabase.com:3306)/sql6519232")
	if err != nil {
		panic(err.Error())
	}
	return db
}
