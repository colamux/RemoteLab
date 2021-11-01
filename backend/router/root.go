package router

import (
	"rlab/middleware"

	"github.com/gin-gonic/gin"
)

func Root(e *gin.Engine) {
	e.GET("/", middleware.InitMiddleWares,
		func(c *gin.Context) {
			c.String(200, "root")
		})
}
