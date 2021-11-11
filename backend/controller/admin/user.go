package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	// base.BaseController
}

// /admin/add
func (con UserController) Add(c *gin.Context) {

}

func (con UserController) Edit(c *gin.Context) {
	fmt.Println("abc")
	fmt.Println(c.GetHeader("Authorization"))
	c.String(200, c.Param("abc"))
}
