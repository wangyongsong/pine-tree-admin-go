package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"one-tree-admin-go/handle/sd"
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
	apiGroup := g.Group("/api")

	user(apiGroup)
	sdFn(g)

	return g
}

//user
func user(group *gin.RouterGroup) {
	user := group.Group("/user")
	{
		user.GET("/info")
	}
}

//sd
func sdFn(engine *gin.Engine) {
	svcd := engine.Group("/sd")

	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}
}
