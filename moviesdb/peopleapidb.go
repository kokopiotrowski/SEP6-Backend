package moviesdb

import (
	"strconv"
	swagger "studies/SEP6-Backend/swagger/models"
	"studies/SEP6-Backend/util"
)

const (
	peopleGetByIdURL    = "/person/"
	peoplePopularGetURL = "/person/popular"
	peopleSearchURL     = "/search/person"
)

func PersonGet(query string, page int64) (swagger.ReturnPeople, error) {
	var returnPeople swagger.ReturnPeople

	params := make(map[string]string)

	params["api_key"] = movieDbAPIKey
	params["query"] = query
	params["page"] = strconv.FormatInt(page, 10)
	code, _, err := util.SendRequest("GET", baseMovieDbURL, peopleSearchURL, params, nil, nil, nil, &returnPeople)
	if err != nil || code != 200 {
		return swagger.ReturnPeople{}, err
	}

	return returnPeople, nil
}

func PersonPersonIdGet(language string, personId int64) (swagger.Person, error) {
	var returnPerson swagger.Person
	params := make(map[string]string)

	params["api_key"] = movieDbAPIKey
	if len(language) > 0 {
		params["language"] = language
	}
	code, _, err := util.SendRequest("GET", baseMovieDbURL, peopleGetByIdURL+strconv.FormatInt(personId, 10), params, nil, nil, nil, &returnPerson)
	if err != nil || code != 200 {
		return swagger.Person{}, err
	}

	return returnPerson, nil
}

func PersonPopularGet(region string, language string, page int64) (swagger.ReturnPeople, error) {
	var returnPeople swagger.ReturnPeople

	params := make(map[string]string)

	params["api_key"] = movieDbAPIKey
	if len(region) > 0 {
		params["region"] = region
	}
	if len(language) > 0 {
		params["language"] = language
	}
	params["page"] = strconv.FormatInt(page, 10)
	code, _, err := util.SendRequest("GET", baseMovieDbURL, peoplePopularGetURL, params, nil, nil, nil, &returnPeople)
	if err != nil || code != 200 {
		return swagger.ReturnPeople{}, err
	}
	return returnPeople, nil
}
