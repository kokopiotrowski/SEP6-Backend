package db

import "database/sql"

var (
	db *sql.DB
)

func connectDB() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:zJjP8kMHj1LeDwFL@tcp(34.116.182.208:3306)/myDB")

	if err != nil {
		panic(err.Error())
	}
	return db, err

}

func GetDB() (*sql.DB, error) {
	var err error
	if db == nil {
		db, err = connectDB()
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}
