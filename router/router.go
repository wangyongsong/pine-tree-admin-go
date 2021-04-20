package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"one-tree-admin-go/router/middleware"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	//404
	g.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "The incorrect API route. 404")
	})

	// The health check handlers
	api := g.Group("/api")

	user(api)

	return g
}

//user
func user(api *gin.RouterGroup) {
	user := api.Group("/user")
	{
		user.GET("/info")
	}
}
