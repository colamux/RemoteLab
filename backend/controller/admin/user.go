package admin

import (
	"fmt"
	"rlab/common"
	"rlab/model/user"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	// base.BaseController
}

// /admin/add
func (con UserController) Add(c *gin.Context) {

	name := c.DefaultPostForm("Name", "cola")

	token, err := common.ReleaseToken(user.User{
		Name:   name,
		Class:  "1",
		Number: "21932032",
	})

	if err != nil {
		c.Abort()
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{"Authorization": token},
		"msg":  name + "登陆成功",
	})
}

func (con UserController) Edit(c *gin.Context) {
	fmt.Println("abc")
	fmt.Println(c.GetHeader("x-token"))
	c.String(200, c.Param("abc"))
}
