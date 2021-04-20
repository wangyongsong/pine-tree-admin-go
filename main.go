package main

import (
	"github.com/gin-gonic/gin"
	"one-tree-admin-go/router"
)

func main() {
	g := gin.Default()

	var middlewares []gin.HandlerFunc

	router.Load(
		g,
		middlewares...,
	)

	_ = g.Run("localhost:9000")
}
