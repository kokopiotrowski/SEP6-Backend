package moviesdb

import (
	"strconv"
	swagger "studies/SEP6-Backend/swagger/models"
	"studies/SEP6-Backend/util"
)

type Cast struct {
	People []swagger.Person `json:cast,omitempty`
}

const (
	movieCastURL = "/movie/"
)

func GetCastForMovie(movieId int64) (Cast, error) {

	var cast Cast

	params := make(map[string]string)

	params["api_key"] = movieDbAPIKey
	code, _, err := util.SendRequest("GET", baseMovieDbURL, movieCastURL+strconv.FormatInt(movieId, 10)+"/credits", params, nil, nil, nil, &cast)
	if err != nil || code != 200 {
		return Cast{}, err
	}

	return cast, nil
}

func GetSimilarMovies(movieId int64) (swagger.ReturnMovies, error) {

	var similarMovies swagger.ReturnMovies

	params := make(map[string]string)

	params["api_key"] = movieDbAPIKey
	code, _, err := util.SendRequest("GET", baseMovieDbURL, movieCastURL+strconv.FormatInt(movieId, 10)+"/similar", params, nil, nil, nil, &similarMovies)
	if err != nil || code != 200 {
		return swagger.ReturnMovies{}, err
	}

	return similarMovies, nil
}
