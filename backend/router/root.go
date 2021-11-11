package router

import (
	"rlab/controller"
	"rlab/middleware"

	"github.com/gin-gonic/gin"
)

func Root(e *gin.Engine) {
	e.GET("/", middleware.InitMiddleWares, middleware.CasbinMiddleware(),
		func(c *gin.Context) {
			c.String(200, "root")
		})

	e.POST("/login", controller.Login)
}
