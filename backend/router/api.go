package router

import (
	"rlab/middleware"

	"github.com/gin-gonic/gin"
)

func Api(e *gin.Engine) {

	apiRouters := e.Group("/api", middleware.JWTMiddleware())

	apiRouters.GET("/", func(c *gin.Context) {
		c.String(200, "api")
	})

	apiRouters.GET("/ls", func(c *gin.Context) {
		c.String(200, "api, ls")
	})
}
