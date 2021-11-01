package router

import (
	"rlab/controller/admin"
	"rlab/middleware"

	"github.com/gin-gonic/gin"
)

func Admin(e *gin.Engine) {

	adminRouters := e.Group("/admin", middleware.InitMiddleWares)

	adminRouters.GET("/", func(c *gin.Context) {
		c.String(200, "admin")
	})

	adminRouters.GET("/user", func(c *gin.Context) {
		c.String(200, "admin, user")
	})

	adminRouters.GET("/add", admin.UserController{}.Add)
}
