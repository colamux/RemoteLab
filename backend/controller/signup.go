package controller

import (
	"rlab/DAL"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	data := make(map[string]string)
	c.BindJSON(&data)

	// 查重

	// 先创建用户、 get uid
	// g
	// if _, err := DAL.Db.

	if _, err := DAL.SetAuth(2, "number", data["username"], data["password"]); err != nil {
		c.JSON(502, gin.H{
			"result": "Signup error",
		})
	}

	c.JSON(200, gin.H{
		"result": "OK",
	})

}
