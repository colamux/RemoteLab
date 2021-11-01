package middleware

import "github.com/gin-gonic/gin"

type jwt struct{}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.
	}
}
