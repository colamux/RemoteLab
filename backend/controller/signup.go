package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	data := make(map[string]string)
	c.BindJSON(&data)

	fmt.Printf("%v", data)

	c.String(200, "注册成功！")

}
