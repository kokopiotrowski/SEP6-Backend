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

func GetCastForMovie(movie *swagger.Movie) (Cast, error) {

	var cast Cast

	var params map[string]string

	params["api_key"] = movieDbAPIKey
	code, _, err := util.SendRequest("GET", baseMovieDbURL, movieCastURL+strconv.FormatInt(movie.Id, 10)+"/credits", params, nil, nil, nil, &cast)
	if err != nil || code != 200 {
		return Cast{}, err
	}

	return cast, nil
}
