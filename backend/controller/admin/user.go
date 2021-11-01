package admin

import (
	"rlab/model/base"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	base.BaseController
}

func (con UserController) Add(c *gin.Context) {
	con.Success(c)
}

func (con UserController) Edit(c *gin.Context) {
	con.Error(c)
}
