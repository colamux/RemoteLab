package controller

import (
	"rlab/DAL"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	data := make(map[string]string)
	c.BindJSON(&data)

	// 查重
	var repeate []int
	if err := DAL.CheckUser(&repeate, data["username"]); err != nil {
		c.JSON(502, gin.H{
			"result": "Signup error",
		})
	}
	if repeate[0] > 0 {
		c.JSON(200, gin.H{
			"result": "用户名重复",
		})
	}

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
