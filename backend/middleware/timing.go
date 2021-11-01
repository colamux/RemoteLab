package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

// 记录访问耗时
func InitMiddleWares(c *gin.Context) {
	start := time.Now().UnixNano()

	c.Next()

	c.String(200, "\n耗时：%d", time.Now().UnixNano()-start)
}
