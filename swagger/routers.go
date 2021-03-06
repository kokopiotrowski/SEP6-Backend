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
	"net/http"
	"strings"
	"studies/SEP6-Backend/util"

	"github.com/gorilla/mux"
)

const (
	IndexMESSAGE = "Hello, this is index page of SEP6 backend app. Try out pinging endpoints. All documentation for this API available here: https://app.swaggerhub.com/apis-docs/k0k0piotrowski/SEP6-Backend/1.0"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {

	util.RespondWithJSON(w, r, http.StatusOK, IndexMESSAGE, nil)
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/k0k0piotrowski/SEP6-Backend/1.0/",
		Index,
	},

	Route{
		"MovieGet",
		strings.ToUpper("Get"),
		"/k0k0piotrowski/SEP6-Backend/1.0/movies",
		MovieGet,
	},

	Route{
		"MovieMovieIdGet",
		strings.ToUpper("Get"),
		"/k0k0piotrowski/SEP6-Backend/1.0/movie/{movieId}",
		MovieMovieIdGet,
	},

	Route{
		"MoviePopularGet",
		strings.ToUpper("Get"),
		"/k0k0piotrowski/SEP6-Backend/1.0/movies/popular",
		MoviePopularGet,
	},

	Route{
		"MovieTopGet",
		strings.ToUpper("Get"),
		"/k0k0piotrowski/SEP6-Backend/1.0/movies/top",
		MovieTopGet,
	},

	Route{
		"PersonGet",
		strings.ToUpper("Get"),
		"/k0k0piotrowski/SEP6-Backend/1.0/people",
		PersonGet,
	},

	Route{
		"PersonPersonIdGet",
		strings.ToUpper("Get"),
		"/k0k0piotrowski/SEP6-Backend/1.0/person/{personId}",
		PersonPersonIdGet,
	},

	Route{
		"PersonPopularGet",
		strings.ToUpper("Get"),
		"/k0k0piotrowski/SEP6-Backend/1.0/people/popular",
		PersonPopularGet,
	},

	Route{
		"UserPlaylistAddToFavouritePost",
		strings.ToUpper("Post"),
		"/k0k0piotrowski/SEP6-Backend/1.0/user/playlist/addToFavourite",
		UserPlaylistAddToFavouritePost,
	},

	Route{
		"UserPlaylistGetFavouriteGet",
		strings.ToUpper("Get"),
		"/k0k0piotrowski/SEP6-Backend/1.0/user/playlist/getFavourite",
		UserPlaylistGetFavouriteGet,
	},

	Route{
		"UserPlaylistRemoveFromFavouriteMovieIdDelete",
		strings.ToUpper("Delete"),
		"/k0k0piotrowski/SEP6-Backend/1.0/user/playlist/removeFromFavourite/{movieId}",
		UserPlaylistRemoveFromFavouriteMovieIdDelete,
	},

	Route{
		"UserLoginPost",
		strings.ToUpper("Post"),
		"/k0k0piotrowski/SEP6-Backend/1.0/user/login",
		UserLoginPost,
	},

	Route{
		"UserRegisterPost",
		strings.ToUpper("Post"),
		"/k0k0piotrowski/SEP6-Backend/1.0/user/register",
		UserRegisterPost,
	},
}
