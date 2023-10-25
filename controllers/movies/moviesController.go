package movies

import (
	templateError "backend_proj_os/errors"
	movieService "backend_proj_os/services/movies"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMovie(c *gin.Context) {
	var request struct {
		ID int `json:"id"`
	}
	if err := c.BindJSON(&request); err != nil {
		statusCode, errResponse := templateError.GetErrorResponse(err)
		c.AbortWithStatusJSON(statusCode, errResponse)
		return
	}
	if request.ID == 0 {
		statusCode, errResponse := templateError.GetErrorResponse(templateError.MissingBodyRequest)
		c.AbortWithStatusJSON(statusCode, errResponse)
		return
	}
	if response, err := movieService.GetMovie(request.ID); err != nil {
		statusCode, errResponse := templateError.GetErrorResponse(err)
		c.AbortWithStatusJSON(statusCode, errResponse)
		return
	} else {
		c.JSON(http.StatusOK, response)
		return
	}
}

func GetMovies(c *gin.Context) {
	if response, err := movieService.GetMovies(); err != nil {
		statusCode, errResponse := templateError.GetErrorResponse(err)
		c.AbortWithStatusJSON(statusCode, errResponse)
		return
	} else {
		c.JSON(http.StatusOK, response)
	}
}

func GetMoviesByTitle(c *gin.Context) {
	var request struct {
		Title string `json:"title"`
	}
	err := c.BindJSON(&request)
	if err != nil {
		statusCode, errResponse := templateError.GetErrorResponse(templateError.InvalidBodyRequest)
		c.AbortWithStatusJSON(statusCode, errResponse)
		return
	}
	if request.Title == "" {
		statusCode, errResponse := templateError.GetErrorResponse(templateError.MissingBodyRequest)
		c.AbortWithStatusJSON(statusCode, errResponse)
		return
	}
	if response, err := movieService.GetMoviesByTitle(request.Title); err != nil {
		statusCode, errResponse := templateError.GetErrorResponse(err)
		c.AbortWithStatusJSON(statusCode, errResponse)
		return
	} else {
		c.JSON(http.StatusOK, response)
	}
}

func GetTopMovies(c *gin.Context) {
	var request struct {
		Count int `json:"count"`
	}
	err := c.BindJSON(&request)
	if err != nil {
		statusCode, errResponse := templateError.GetErrorResponse(templateError.InvalidBodyRequest)
		c.AbortWithStatusJSON(statusCode, errResponse)
		return
	}
	if request.Count == 0 {
		statusCode, errResponse := templateError.GetErrorResponse(templateError.MissingBodyRequest)
		c.AbortWithStatusJSON(statusCode, errResponse)
		return
	}

	if response, err := movieService.GetTopMovies(request.Count); err != nil {
		statusCode, errResponse := templateError.GetErrorResponse(err)
		c.AbortWithStatusJSON(statusCode, errResponse)
		return
	} else {
		c.JSON(http.StatusOK, response)
	}
}

func GetMoviesByCategory(c *gin.Context) {
	var request struct {
		Category string `json:"category"`
	}
	err := c.BindJSON(&request)
	if err != nil {
		statusCode, errResponse := templateError.GetErrorResponse(templateError.InvalidBodyRequest)
		c.AbortWithStatusJSON(statusCode, errResponse)
		return
	}
	if request.Category == "" {
		statusCode, errResponse := templateError.GetErrorResponse(templateError.MissingBodyRequest)
		c.AbortWithStatusJSON(statusCode, errResponse)
		return
	}
	if response, err := movieService.GetMoviesByCategory(request.Category); err != nil {
		statusCode, errResponse := templateError.GetErrorResponse(err)
		c.AbortWithStatusJSON(statusCode, errResponse)
		return
	} else {
		c.JSON(http.StatusOK, response)
	}
}

func GetMoviesByDirectorName(c *gin.Context) {
	var request struct {
		Director string `json:"directorName"`
	}
	err := c.BindJSON(&request)
	if err != nil {
		statusCode, errResponse := templateError.GetErrorResponse(templateError.InvalidBodyRequest)
		c.AbortWithStatusJSON(statusCode, errResponse)
		return
	}
	if request.Director == "" {
		statusCode, errResponse := templateError.GetErrorResponse(templateError.MissingBodyRequest)
		c.AbortWithStatusJSON(statusCode, errResponse)
		return
	}
	if response, err := movieService.GetMoviesByDirector(request.Director); err != nil {
		statusCode, errResponse := templateError.GetErrorResponse(err)
		c.AbortWithStatusJSON(statusCode, errResponse)
		return
	} else {
		c.JSON(http.StatusOK, response)
	}
}
