package moviesdb

import (
	"strconv"
	swagger "studies/SEP6-Backend/swagger/models"
	"studies/SEP6-Backend/util"
)

const (
	moviesGetPopularURL = "/movie/popular"
	moviesSearchURL     = "/search/movie"
	moviesGetByIdURL    = "/movie/" //{movie_id}
	moviesGetTopURL     = "/movie/top_rated"
)

func MoviesGetPopular(region string, language string, page int64) (swagger.ReturnMovies, error) {

	var returnMovies swagger.ReturnMovies

	params := make(map[string]string)

	params["api_key"] = movieDbAPIKey
	if len(region) > 0 {
		params["region"] = region
	}
	if len(language) > 0 {
		params["language"] = language
	}
	params["page"] = strconv.FormatInt(page, 10)
	code, _, err := util.SendRequest("GET", baseMovieDbURL, moviesGetPopularURL, params, nil, nil, nil, &returnMovies)
	if err != nil || code != 200 {
		return swagger.ReturnMovies{}, err
	}
	return returnMovies, nil
}

func MovieSearch(query string, page int64) (swagger.ReturnMovies, error) {

	var returnMovies swagger.ReturnMovies

	params := make(map[string]string)

	params["api_key"] = movieDbAPIKey
	params["query"] = query
	params["page"] = strconv.FormatInt(page, 10)
	code, _, err := util.SendRequest("GET", baseMovieDbURL, moviesSearchURL, params, nil, nil, nil, &returnMovies)
	if err != nil || code != 200 {
		return swagger.ReturnMovies{}, err
	}

	return returnMovies, nil
}

func MovieMovieIdGet(language string, movieId int64) (swagger.Movie, error) {

	var returnMovie swagger.Movie
	params := make(map[string]string)

	params["api_key"] = movieDbAPIKey
	if len(language) > 0 {
		params["language"] = language
	}
	code, _, err := util.SendRequest("GET", baseMovieDbURL, moviesGetByIdURL+strconv.FormatInt(movieId, 10), params, nil, nil, nil, &returnMovie)
	if err != nil || code != 200 {
		return swagger.Movie{}, err
	}
	similarMovies, err := GetSimilarMovies(returnMovie.Id)
	if err != nil {
		return swagger.Movie{}, err
	}
	cast, err := GetCastForMovie(returnMovie.Id)
	if err != nil {
		return swagger.Movie{}, err
	}

	returnMovie.Cast = cast.People
	returnMovie.SimilarMovies = similarMovies.Movies

	return returnMovie, nil
}

func MovieTopGet(region string, language string, page int64) (swagger.ReturnMovies, error) {

	var returnMovies swagger.ReturnMovies

	params := make(map[string]string)

	params["api_key"] = movieDbAPIKey
	if len(region) > 0 {
		params["region"] = region
	}
	if len(language) > 0 {
		params["language"] = language
	}
	params["page"] = strconv.FormatInt(page, 10)
	code, _, err := util.SendRequest("GET", baseMovieDbURL, moviesGetTopURL, params, nil, nil, nil, &returnMovies)
	if err != nil || code != 200 {
		return swagger.ReturnMovies{}, err
	}
	return returnMovies, nil
}
