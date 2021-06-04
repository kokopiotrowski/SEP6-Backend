/*
 * SEP6-movies backend
 *
 * Backend part of the project delivered for SEP6 course - Movies platform  Authors of project:  Konrad Piotrowski (280053) Aleksander Stefan Bialik (280027)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"studies/SEP6-Backend/moviesdb"
	"studies/SEP6-Backend/util"
)

func PersonGet(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	search, present := q["search"]
	if !present || len(search) == 0 {
		search = []string{""}
	}

	page, present := q["page"]
	if !present || len(page) == 0 {
		page = []string{"1"}
	}

	pageParsed, err := strconv.ParseInt(page[0], 10, 64)
	if err != nil {
		util.RespondWithJSON(w, r, http.StatusInternalServerError, nil, errors.New("Failed to parse pagination"))
		return
	}
	searchedMovies, err := moviesdb.PersonGet(strings.Join(search, ""), pageParsed)
	if err != nil {
		util.RespondWithJSON(w, r, http.StatusInternalServerError, nil, errors.New("Failed to search for people"))
		return
	}
	util.RespondWithJSON(w, r, http.StatusOK, searchedMovies, nil)
}

func PersonPersonIdGet(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var personId int64

	if util.DecodeVarAsInt64(w, r, "personId", &personId) {

		language, present := q["language"]
		if !present || len(language) == 0 {
			language = []string{""}
		}
		receivedPerson, err := moviesdb.PersonPersonIdGet(strings.Join(language, ""), personId)
		if err != nil {
			util.RespondWithJSON(w, r, http.StatusInternalServerError, nil, errors.New("Failed to get person"))
			return
		}
		util.RespondWithJSON(w, r, http.StatusOK, receivedPerson, nil)
	}
}

func PersonPopularGet(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	page, present := q["page"]
	if !present || len(page) == 0 {
		page = []string{"1"}
	}
	pageParsed, err := strconv.ParseInt(strings.Join(page, ""), 10, 64)
	if err != nil {
		util.RespondWithJSON(w, r, http.StatusInternalServerError, nil, errors.New("Failed to parse pagination"))
		return
	}
	language, present := q["language"]
	if !present || len(language) == 0 {
		language = []string{""}
	}
	popularPeople, err := moviesdb.PersonPopularGet("", strings.Join(language, ""), pageParsed)
	if err != nil {
		util.RespondWithJSON(w, r, http.StatusInternalServerError, nil, errors.New("Failed to get popular people"))
		return
	}
	util.RespondWithJSON(w, r, http.StatusOK, popularPeople, nil)
}
