package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	swagger "studies/SEP6-Backend/swagger/models"

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

func AddFavouriteMovie(userId int64, movieId int64, title string, poster_path string) error {

	query := "INSERT INTO FavouriteMovies(UserId, MovieId, Title, PosterPath) VALUES (?, ?, ?, ?)"

	tx, err := db.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec(query, userId, movieId, title, poster_path)

	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func GetFavouriteMovies(userId int64) ([]swagger.FavouriteMovie, error) {

	query := "SELECT FavouriteMovies.MovieId, FavouriteMovies.Title, FavouriteMovies.PosterPath FROM FavouriteMovies WHERE UserId=?"

	tx, err := db.Begin()

	if err != nil {
		return []swagger.FavouriteMovie{}, err
	}

	results, err := tx.Query(query, userId)

	if err != nil {
		return []swagger.FavouriteMovie{}, err
	}
	defer tx.Commit()
	var favouriteMovies []swagger.FavouriteMovie
	var movieId int64
	var title string
	var poster_path string

	for results.Next() {
		err = results.Scan(&movieId, &title, &poster_path)
		if err != nil {
			return []swagger.FavouriteMovie{}, err
		}
		fm := swagger.FavouriteMovie{
			MovieId:    movieId,
			Title:      title,
			PosterPath: poster_path,
		}
		favouriteMovies = append(favouriteMovies, fm)
	}
	return favouriteMovies, nil
}

func DeleteFavouriteMovie(userId int64, movieId int64) error {

	query := "DELETE FROM FavouriteMovies WHERE UserId=? AND MovieId=?"

	tx, err := db.Begin()

	if err != nil {
		return err
	}

	result, err := tx.Exec(query, userId, movieId)

	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("No such movie with this ID to remove from playlist")
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
