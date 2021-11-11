package controller

import (
	"fmt"
	"rlab/DAL"
	"rlab/common"
	"rlab/model/user"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	// c.PostForm("username") // form-data格式，选用json格式
	data := make(map[string]string)
	c.BindJSON(&data)
	var name string = data["username"]

	// var password []string
	fmt.Println(data["password"])
	// err := DAL.GetAuth(&password, "number", name)
	passwd, err := DAL.GetPasswd("number", name)
	// fmt.Println(password[0])
	if err != nil {
		c.String(500, "服务器开小差了~~")
		return
	}
	if data["password"] != passwd {
		c.String(401, "用户名or密码错误")
		return
	}
	// fmt.Println(data["password"])

	// get vm
	var vms []string
	fmt.Println(name)
	err = DAL.GetVM(&vms, name)
	if err != nil {
		fmt.Errorf("sql query err: %s", err.Error())
	}
	for _, u := range vms {
		fmt.Println(u)
	}

	token, err := common.ReleaseToken(user.UserInfo{
		Name:   name,
		Class:  1,
		Number: 21932032,
		Id:     1,
		Vm:     "",
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
