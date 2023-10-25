package test

import (
	"net/http"

	templateError "backend_proj_os/errors"
	testService "backend_proj_os/services/test"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}

func TestModel(c *gin.Context) {
	if response, err := testService.GetAllData(); err != nil {
		statusCode, errResponse := templateError.GetErrorResponse(err)
		c.AbortWithStatusJSON(statusCode, errResponse)

	} else {
		c.JSON(http.StatusOK, response)
	}
}
