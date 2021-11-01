package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleWares(c *gin.Context) {
	start := time.Now().UnixNano()

	c.Next()

	c.String(200, "\n耗时：%d", time.Now().UnixNano()-start)
}
