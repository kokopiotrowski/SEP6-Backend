package moviesdb

import (
	"strconv"
	swagger "studies/SEP6-Backend/swagger/models"
	"studies/SEP6-Backend/util"
)

type PlayingPeople struct {
	People []swagger.Person `json:"cast,omitempty"`
}

type CastMovies struct {
	Movies []swagger.Movie `json:"cast,omitempty"`
}

const (
	movieCastURL = "/movie/"
)

func GetCastForMovie(movieId int64, language string) (PlayingPeople, error) {

	var cast PlayingPeople

	params := make(map[string]string)

	params["api_key"] = movieDbAPIKey
	if len(language) > 0 {
		params["language"] = language
	}
	code, _, err := util.SendRequest("GET", baseMovieDbURL, movieCastURL+strconv.FormatInt(movieId, 10)+"/credits", params, nil, nil, nil, &cast)
	if err != nil || code != 200 {
		return PlayingPeople{}, err
	}

	return cast, nil
}

func GetSimilarMovies(movieId int64, language string) (swagger.ReturnMovies, error) {

	var similarMovies swagger.ReturnMovies

	params := make(map[string]string)
	params["api_key"] = movieDbAPIKey
	if len(language) > 0 {
		params["language"] = language
	}
	code, _, err := util.SendRequest("GET", baseMovieDbURL, movieCastURL+strconv.FormatInt(movieId, 10)+"/similar", params, nil, nil, nil, &similarMovies)
	if err != nil || code != 200 {
		return swagger.ReturnMovies{}, err
	}

	return similarMovies, nil
}

func GetMoviesByPersonId(personId int64) (CastMovies, error) {

	var movies CastMovies

	params := make(map[string]string)

	params["api_key"] = movieDbAPIKey
	code, _, err := util.SendRequest("GET", baseMovieDbURL, "/person/"+strconv.FormatInt(personId, 10)+"/movie_credits", params, nil, nil, nil, &movies)
	if err != nil || code != 200 {
		return CastMovies{}, err
	}

	return movies, nil
}
