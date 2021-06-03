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
	"fmt"
	"net/http"
	"strconv"
	"studies/SEP6-Backend/moviesdb"
	swagger "studies/SEP6-Backend/swagger/models"
	"studies/SEP6-Backend/util"
)

var (
	dummyMovies swagger.ReturnMovies = swagger.ReturnMovies{
		Page:         1,
		TotalPages:   3,
		TotalResults: 4,
		Movies: []swagger.Movie{
			{
				Id:            1,
				PosterPath:    "jakiś tam path",
				Title:         "Wariat",
				Cast:          []swagger.Person{},
				VoteAverage:   1.2,
				VoteCount:     3,
				SimilarMovies: []swagger.Movie{},
			},
			{
				Id:            2,
				PosterPath:    "jakiś tam path numer 2",
				Title:         "Elo",
				Cast:          []swagger.Person{},
				VoteAverage:   1.2,
				VoteCount:     2,
				SimilarMovies: []swagger.Movie{},
			},
			{
				Id:            3,
				PosterPath:    "jakiś tam path sadasd21",
				Title:         "Siema",
				Cast:          []swagger.Person{},
				VoteAverage:   1.2,
				VoteCount:     5,
				SimilarMovies: []swagger.Movie{},
			},
			{
				Id:            4,
				PosterPath:    "jakiś tam path sadasd1212312312312312312",
				Title:         "hmmmm",
				Cast:          []swagger.Person{},
				VoteAverage:   2.3,
				VoteCount:     4,
				SimilarMovies: []swagger.Movie{},
			},
		},
	}
)

func MovieGet(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	search, present := q["search"]
	if !present || len(search[0]) == 0 {
		fmt.Println("region is not present")
	}

	page, present := q["page"]
	if !present || len(page[0]) == 0 {
		fmt.Println("page not present")
	}

	pageParsed, err := strconv.ParseInt(page[0], 10, 64)
	if err != nil {
		util.RespondWithJSON(w, r, http.StatusInternalServerError, nil, errors.New("Failed to parse pagination"))
		return
	}
	searchedMovies, err := moviesdb.MovieSearch(search[0], pageParsed)
	if err != nil {
		util.RespondWithJSON(w, r, http.StatusInternalServerError, nil, errors.New("Failed to search for movies"))
		return
	}
	util.RespondWithJSON(w, r, http.StatusOK, searchedMovies, nil)
}

func MovieMovieIdGet(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var movieId int64

	if util.DecodeVarAsInt64(w, r, "movieId", &movieId) {

		language, present := q["language"]
		if !present || len(language[0]) == 0 {
			fmt.Println("language not present")
		}
		searchedMovies, err := moviesdb.MovieMovieIdGet(language[0], movieId)
		if err != nil {
			util.RespondWithJSON(w, r, http.StatusInternalServerError, nil, errors.New("Failed to search for movies"))
			return
		}
		util.RespondWithJSON(w, r, http.StatusOK, searchedMovies, nil)
	}

}

func MoviePopularGet(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	page, present := q["page"]
	if !present || len(page[0]) == 0 {
		fmt.Println("page not present")
	}

	pageParsed, err := strconv.ParseInt(page[0], 10, 64)
	if err != nil {
		util.RespondWithJSON(w, r, http.StatusInternalServerError, nil, errors.New("Failed to parse pagination"))
		return
	}
	popularMovies, err := moviesdb.MoviesGetPopular("", "", pageParsed)
	if err != nil {
		util.RespondWithJSON(w, r, http.StatusInternalServerError, nil, errors.New("Failed to search for movies"))
		return
	}
	util.RespondWithJSON(w, r, http.StatusOK, popularMovies, nil)
}

func MovieTopGet(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	page, present := q["page"]
	if !present || len(page[0]) == 0 {
		fmt.Println("page not present")
	}

	pageParsed, err := strconv.ParseInt(page[0], 10, 64)
	if err != nil {
		util.RespondWithJSON(w, r, http.StatusInternalServerError, nil, errors.New("Failed to parse pagination"))
		return
	}
	topMovies, err := moviesdb.MovieTopGet("", "", pageParsed)
	if err != nil {
		util.RespondWithJSON(w, r, http.StatusInternalServerError, nil, errors.New("Failed to search for movies"))
		return
	}
	util.RespondWithJSON(w, r, http.StatusOK, topMovies, nil)
}
