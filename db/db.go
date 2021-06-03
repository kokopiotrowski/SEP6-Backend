package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *sql.DB
)

func connectDB() (*sql.DB, error) {

	dsn := "root:1s2dpLJivs3Cvla3@tcp(34.116.246.37:3306)/myDB?charset=utf8mb4&parseTime=True&loc=Local"

	// dbPool is the pool of database connections.
	dbGorm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}
	db, err = dbGorm.DB()
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

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}
