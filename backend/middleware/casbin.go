package middleware

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		e, _ := casbin.NewEnforcer("config/model.conf", "config/policy.csv")

		//定义各种sub，obj和act的数组
		subs := []string{"bob", "zeta"}
		objs := []string{"data1", "data2"}
		acts := []string{"read", "write"}

		//遍历组合sub，obj，act并打印出对应策略匹配结果。
		for _, sub := range subs {
			for _, obj := range objs {
				for _, act := range acts {
					v, _ := e.Enforce(sub, obj, act)
					fmt.Println(sub, "/", obj, "/", act, "=", v)
				}
			}
		}

		fmt.Println(e.GetFilteredPolicy(0, "bob"))
	}
}
