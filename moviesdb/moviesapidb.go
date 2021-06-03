package moviesdb

import (
	swagger "studies/SEP6-Backend/swagger/models"
	"studies/SEP6-Backend/util"
)

func MoviesGetPopular() ([]swagger.Movie, error) {

	util.SendRequest()

	return nil, nil
}

func MovieSearch() (swagger.ReturnMovies, error) {

	return swagger.ReturnMovies{}, nil
}

func MovieMovieIdGet() (swagger.Movie, error) {

	return swagger.Movie{}, nil
}

func MovieTopGet() (swagger.ReturnMovies, error) {

	return swagger.ReturnMovies{}, nil
}
