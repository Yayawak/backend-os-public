package movies

import (
	moviesController "backend_proj_os/controllers/movies"
	// testController "backend_proj_os/controllers/test"

	"github.com/gin-gonic/gin"
)

func MoviesRouter(router *gin.Engine) {
	paths := "/movies"
	path := "/movie"
	router.GET(path+"/search", moviesController.GetMovie)
	router.GET(paths+"/search", moviesController.GetMovies)
	router.GET(paths+"/search/title", moviesController.GetMoviesByTitle)
	router.GET(paths+"/top", moviesController.GetTopMovies)
	router.GET(paths+"/search/category", moviesController.GetMoviesByCategory)
	router.GET(paths+"/search/director", moviesController.GetMoviesByDirectorName)
}
