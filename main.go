package main

import (
	moviesRouter "backend_proj_os/routes/movies"
	testRouter "backend_proj_os/routes/test"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	routes := gin.New()

	routes.Use(cors.Default())
	testRouter.TestRoute(routes)
	moviesRouter.MoviesRouter(routes)

	routes.Run() // default is port : 8080 , can set port by routes.Run(":3000")
}
