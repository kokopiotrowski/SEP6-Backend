package db

import "database/sql"

func connectDB() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:zJjP8kMHj1LeDwFL@tcp(34.116.182.208:3306)/myDB")

	if err != nil {
		panic(err.Error())
	}
	return db, err

}
